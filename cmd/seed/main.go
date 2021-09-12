package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/Netflix/go-env"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/coursegrade"
	"github.com/msal4/hassah_school_server/ent/user"
	"github.com/msal4/hassah_school_server/server"
	"github.com/msal4/hassah_school_server/server/model"
	"github.com/msal4/hassah_school_server/service"
	"github.com/msal4/hassah_school_server/util/ptr"
	"gopkg.in/yaml.v2"
)

func main() {
	f, err := os.Open("./config.yml")
	if err != nil {
		log.Fatal(err)
	}

	var cfg server.Config
	err = yaml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	godotenv.Load()

	env.UnmarshalFromEnviron(&cfg)

	if cfg.Port == 0 {
		cfg.Port = 3000
	}

	if cfg.DatabaseDialect == "" {
		cfg.DatabaseDialect = cfg.DatabaseURL[:strings.Index(cfg.DatabaseURL, ":")]
	}

	ec, err := ent.Open(cfg.DatabaseDialect, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed establishing connection: %v", err)
	}
	defer ec.Close()

	mc, err := minio.New(cfg.Minio.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(cfg.Minio.AccessKey, cfg.Minio.Token, ""),
	})
	if err != nil {
		log.Fatalf("instantiating minio client: %v", err)
	}

	s, err := service.New(ec, mc, &cfg.Config)
	if err != nil {
		log.Fatalf("creating service: %v", err)
	}

	if err = seed(context.Background(), s); err != nil {
		log.Fatalf("seeding: %v", err)
	}

	log.Println("Finished seeding 🎉.")
}

func seed(ctx context.Context, s *service.Service) error {
	f, err := os.Open("./testfiles/harvard.jpg")
	if err != nil {
		return err
	}

	stat, err := f.Stat()
	if err != nil {
		return err
	}

	sch, err := s.AddSchool(ctx, model.AddSchoolInput{
		Name:   "Palestine School",
		Active: true,
		Image: graphql.Upload{
			File:        f,
			Filename:    f.Name(),
			Size:        stat.Size(),
			ContentType: "image/jpg",
		},
	})
	if err != nil {
		return err
	}
	log.Printf("Created school: %v\n\n", sch)

	stg, err := s.AddStage(ctx, model.AddStageInput{
		Name:          "1st Elementry Stage",
		Active:        true,
		TuitionAmount: 15000000,
		SchoolID:      sch.ID,
	})
	if err != nil {
		return err
	}
	log.Printf("Created stage: %v\n\n", stg)

	tchr, err := s.AddUser(ctx, model.AddUserInput{
		Name:     "John Doe Teacher",
		Username: "teacher01",
		Password: "teacher01pass",
		Phone:    "07712345672",
		Role:     user.RoleTeacher,
		SchoolID: &sch.ID,
		Active:   true,
	})
	if err != nil {
		return err
	}
	log.Printf("Created teacher: %v\n\n", tchr)

	tchr2, err := s.AddUser(ctx, model.AddUserInput{
		Name:     "Jane Doe Teacher",
		Username: "teacher02",
		Password: "teacher02pass",
		Phone:    "07712345673",
		Role:     user.RoleTeacher,
		SchoolID: &sch.ID,
		Active:   true,
	})
	if err != nil {
		return err
	}
	log.Printf("Created teacher: %v\n\n", tchr2)

	cls, err := s.AddClass(ctx, model.AddClassInput{
		Name:      "الرياضيات",
		Active:    true,
		StageID:   stg.ID,
		TeacherID: tchr.ID,
	})
	log.Printf("Created class: %v\n\n", cls)
	if err != nil {
		return err
	}

	grp, err := cls.Group(ctx)
	if err != nil {
		return err
	}

	cls2, err := s.AddClass(ctx, model.AddClassInput{
		Name:      "الفيزياء",
		Active:    true,
		StageID:   stg.ID,
		TeacherID: tchr2.ID,
	})
	log.Printf("Created class: %v\n\n", cls2)
	if err != nil {
		return err
	}

	cls3, err := s.AddClass(ctx, model.AddClassInput{
		Name:      "English",
		Active:    true,
		StageID:   stg.ID,
		TeacherID: tchr.ID,
	})
	if err != nil {
		return err
	}
	log.Printf("Created class: %v\n\n", cls3)

	scd, err := s.AddSchedule(ctx, model.AddScheduleInput{ClassID: cls.ID, Weekday: time.Sunday, Duration: time.Hour / 2, StartsAt: time.Now()})
	if err != nil {
		return err
	}
	log.Printf("Created schedule: %v\n\n", scd)

	scd, err = s.AddSchedule(ctx, model.AddScheduleInput{ClassID: cls2.ID, Weekday: time.Sunday, Duration: time.Hour / 3, StartsAt: time.Now().Add(time.Hour)})
	if err != nil {
		return err
	}
	log.Printf("Created schedule: %v\n\n", scd)

	scd, err = s.AddSchedule(ctx, model.AddScheduleInput{ClassID: cls2.ID, Weekday: time.Monday, Duration: time.Hour / 3, StartsAt: time.Now().Add(-time.Hour)})
	if err != nil {
		return err
	}
	log.Printf("Created schedule: %v\n\n", scd)

	scd, err = s.AddSchedule(ctx, model.AddScheduleInput{ClassID: cls2.ID, Weekday: time.Tuesday, Duration: time.Hour / 3, StartsAt: time.Now().Add(-time.Hour / 2)})
	if err != nil {
		return err
	}
	log.Printf("Created schedule: %v\n\n", scd)

	scd, err = s.AddSchedule(ctx, model.AddScheduleInput{ClassID: cls3.ID, Weekday: time.Tuesday, Duration: time.Hour / 2, StartsAt: time.Now().Add(-time.Hour / 3)})
	if err != nil {
		return err
	}
	log.Printf("Created schedule: %v\n\n", scd)

	scd, err = s.AddSchedule(ctx, model.AddScheduleInput{ClassID: cls3.ID, Weekday: time.Monday, Duration: time.Hour / 2, StartsAt: time.Now().Add(time.Hour / 2)})
	if err != nil {
		return err
	}
	log.Printf("Created schedule: %v\n\n", scd)

	f.Seek(0, 0)
	stdt, err := s.AddUser(ctx, model.AddUserInput{
		Name:     "John Doe",
		Username: "student01",
		Password: "student01pass",
		Phone:    "07712345678",
		Role:     user.RoleStudent,
		StageID:  &stg.ID,
		Active:   true,
		Image: &graphql.Upload{
			File:        f,
			Filename:    f.Name(),
			Size:        stat.Size(),
			ContentType: "image/jpeg",
		},
	})
	if err != nil {
		return err
	}
	log.Printf("Created student: %v\n\n", stdt)

	_, err = s.AddCourseGrade(ctx, model.AddCourseGradeInput{
		Course:         coursegrade.CourseFirst,
		StudentID:      stdt.ID,
		StageID:        stg.ID,
		ClassID:        cls.ID,
		ActivityFirst:  ptr.Int(10),
		ActivitySecond: ptr.Int(20),
		WrittenFirst:   ptr.Int(16),
		WrittenSecond:  ptr.Int(17),
		CourseFinal:    ptr.Int(30),
		Year:           "2020-2021",
	})
	if err != nil {
		return err
	}
	_, err = s.AddCourseGrade(ctx, model.AddCourseGradeInput{
		Course:         coursegrade.CourseFirst,
		StudentID:      stdt.ID,
		StageID:        stg.ID,
		ClassID:        cls.ID,
		ActivityFirst:  ptr.Int(14),
		ActivitySecond: ptr.Int(22),
		WrittenFirst:   ptr.Int(12),
		WrittenSecond:  ptr.Int(19),
		CourseFinal:    ptr.Int(31),
		Year:           "2019-2020",
	})
	if err != nil {
		return err
	}
	_, err = s.AddCourseGrade(ctx, model.AddCourseGradeInput{
		Course:         coursegrade.CourseSecond,
		StudentID:      stdt.ID,
		StageID:        stg.ID,
		ClassID:        cls.ID,
		ActivityFirst:  ptr.Int(10),
		ActivitySecond: ptr.Int(13),
		WrittenFirst:   ptr.Int(19),
		WrittenSecond:  ptr.Int(13),
		CourseFinal:    ptr.Int(38),
		Year:           "2020-2021",
	})
	if err != nil {
		return err
	}

	s.AddGroup(ctx, service.AddGroupInput{UserIDs: []uuid.UUID{stdt.ID, tchr.ID}, Active: true})
	s.AddGroup(ctx, service.AddGroupInput{UserIDs: []uuid.UUID{stdt.ID, tchr2.ID}, Active: true})

	s.EC.Message.Create().SetContent("gibberish content 1").SetGroup(grp).SetOwner(stdt).Exec(ctx)
	s.EC.Message.Create().SetContent("gibberish content 2").SetGroup(grp).SetOwner(stdt).Exec(ctx)
	s.EC.Message.Create().SetContent("gibberish content 3").SetGroup(grp).SetOwner(stdt).Exec(ctx)
	s.EC.Message.Create().SetContent("gibberish content 4").SetGroup(grp).SetOwner(stdt).Exec(ctx)
	s.EC.Message.Create().SetContent("gibberish content 5").SetGroup(grp).SetOwner(stdt).Exec(ctx)
	s.EC.Message.Create().SetContent("gibberish content 6").SetGroup(grp).SetOwner(stdt).Exec(ctx)
	s.EC.Message.Create().SetContent("gibberish content 7").SetGroup(grp).SetOwner(stdt).Exec(ctx)
	s.EC.Message.Create().SetContent("gibberish content 8").SetGroup(grp).SetOwner(stdt).Exec(ctx)

	f.Seek(0, 0)

	ass, err := s.AddAssignment(ctx, model.AddAssignmentInput{
		ClassID:     cls3.ID,
		Name:        "new assign",
		Description: ptr.Str("descsdfksdkfj s"),
		File:        &graphql.Upload{File: f, Filename: f.Name(), Size: stat.Size(), ContentType: "image/jpeg"},
		DueDate:     time.Now().Add(time.Hour * 100),
	})
	if err != nil {
		return err
	}
	log.Printf("Created assignment: %v\n\n", ass)

	f.Seek(0, 0)
	sub, err := s.AddAssignmentSubmission(ctx, stdt.ID, model.AddAssignmentSubmissionInput{
		AssignmentID: ass.ID,
		Files:        []*graphql.Upload{{File: f, Filename: f.Name(), Size: stat.Size(), ContentType: "image/jpeg"}},
	})
	if err != nil {
		return err
	}
	log.Printf("Created submission: %v\n\n", sub)

	payment, err := s.AddTuitionPayment(ctx, model.AddTuitionPaymentInput{
		StageID:    stg.ID,
		StudentID:  stdt.ID,
		PaidAmount: 100000,
		Year:       "2020-2021",
	})
	if err != nil {
		return err
	}
	log.Printf("Created payment: %v\n\n", payment)

	payment, err = s.AddTuitionPayment(ctx, model.AddTuitionPaymentInput{
		StageID:    stg.ID,
		StudentID:  stdt.ID,
		PaidAmount: 50000,
		Year:       "2020-2021",
	})
	if err != nil {
		return err
	}
	log.Printf("Created payment: %v\n\n", payment)

	return nil
}
