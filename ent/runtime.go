// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent/assignment"
	"github.com/msal4/hassah_school_server/ent/assignmentsubmission"
	"github.com/msal4/hassah_school_server/ent/attendance"
	"github.com/msal4/hassah_school_server/ent/class"
	"github.com/msal4/hassah_school_server/ent/coursegrade"
	"github.com/msal4/hassah_school_server/ent/grade"
	"github.com/msal4/hassah_school_server/ent/group"
	"github.com/msal4/hassah_school_server/ent/message"
	"github.com/msal4/hassah_school_server/ent/notification"
	"github.com/msal4/hassah_school_server/ent/schedule"
	"github.com/msal4/hassah_school_server/ent/schema"
	"github.com/msal4/hassah_school_server/ent/school"
	"github.com/msal4/hassah_school_server/ent/stage"
	"github.com/msal4/hassah_school_server/ent/tuitionpayment"
	"github.com/msal4/hassah_school_server/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	assignmentMixin := schema.Assignment{}.Mixin()
	assignmentMixinFields0 := assignmentMixin[0].Fields()
	_ = assignmentMixinFields0
	assignmentFields := schema.Assignment{}.Fields()
	_ = assignmentFields
	// assignmentDescCreatedAt is the schema descriptor for created_at field.
	assignmentDescCreatedAt := assignmentMixinFields0[0].Descriptor()
	// assignment.DefaultCreatedAt holds the default value on creation for the created_at field.
	assignment.DefaultCreatedAt = assignmentDescCreatedAt.Default.(func() time.Time)
	// assignmentDescUpdatedAt is the schema descriptor for updated_at field.
	assignmentDescUpdatedAt := assignmentMixinFields0[1].Descriptor()
	// assignment.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	assignment.DefaultUpdatedAt = assignmentDescUpdatedAt.Default.(func() time.Time)
	// assignment.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	assignment.UpdateDefaultUpdatedAt = assignmentDescUpdatedAt.UpdateDefault.(func() time.Time)
	// assignmentDescName is the schema descriptor for name field.
	assignmentDescName := assignmentFields[1].Descriptor()
	// assignment.NameValidator is a validator for the "name" field. It is called by the builders before save.
	assignment.NameValidator = assignmentDescName.Validators[0].(func(string) error)
	// assignmentDescIsExam is the schema descriptor for is_exam field.
	assignmentDescIsExam := assignmentFields[4].Descriptor()
	// assignment.DefaultIsExam holds the default value on creation for the is_exam field.
	assignment.DefaultIsExam = assignmentDescIsExam.Default.(bool)
	// assignmentDescID is the schema descriptor for id field.
	assignmentDescID := assignmentFields[0].Descriptor()
	// assignment.DefaultID holds the default value on creation for the id field.
	assignment.DefaultID = assignmentDescID.Default.(func() uuid.UUID)
	assignmentsubmissionMixin := schema.AssignmentSubmission{}.Mixin()
	assignmentsubmissionMixinFields0 := assignmentsubmissionMixin[0].Fields()
	_ = assignmentsubmissionMixinFields0
	assignmentsubmissionFields := schema.AssignmentSubmission{}.Fields()
	_ = assignmentsubmissionFields
	// assignmentsubmissionDescCreatedAt is the schema descriptor for created_at field.
	assignmentsubmissionDescCreatedAt := assignmentsubmissionMixinFields0[0].Descriptor()
	// assignmentsubmission.DefaultCreatedAt holds the default value on creation for the created_at field.
	assignmentsubmission.DefaultCreatedAt = assignmentsubmissionDescCreatedAt.Default.(func() time.Time)
	// assignmentsubmissionDescUpdatedAt is the schema descriptor for updated_at field.
	assignmentsubmissionDescUpdatedAt := assignmentsubmissionMixinFields0[1].Descriptor()
	// assignmentsubmission.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	assignmentsubmission.DefaultUpdatedAt = assignmentsubmissionDescUpdatedAt.Default.(func() time.Time)
	// assignmentsubmission.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	assignmentsubmission.UpdateDefaultUpdatedAt = assignmentsubmissionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// assignmentsubmissionDescID is the schema descriptor for id field.
	assignmentsubmissionDescID := assignmentsubmissionFields[0].Descriptor()
	// assignmentsubmission.DefaultID holds the default value on creation for the id field.
	assignmentsubmission.DefaultID = assignmentsubmissionDescID.Default.(func() uuid.UUID)
	attendanceMixin := schema.Attendance{}.Mixin()
	attendanceMixinFields0 := attendanceMixin[0].Fields()
	_ = attendanceMixinFields0
	attendanceFields := schema.Attendance{}.Fields()
	_ = attendanceFields
	// attendanceDescCreatedAt is the schema descriptor for created_at field.
	attendanceDescCreatedAt := attendanceMixinFields0[0].Descriptor()
	// attendance.DefaultCreatedAt holds the default value on creation for the created_at field.
	attendance.DefaultCreatedAt = attendanceDescCreatedAt.Default.(func() time.Time)
	// attendanceDescUpdatedAt is the schema descriptor for updated_at field.
	attendanceDescUpdatedAt := attendanceMixinFields0[1].Descriptor()
	// attendance.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	attendance.DefaultUpdatedAt = attendanceDescUpdatedAt.Default.(func() time.Time)
	// attendance.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	attendance.UpdateDefaultUpdatedAt = attendanceDescUpdatedAt.UpdateDefault.(func() time.Time)
	// attendanceDescID is the schema descriptor for id field.
	attendanceDescID := attendanceFields[0].Descriptor()
	// attendance.DefaultID holds the default value on creation for the id field.
	attendance.DefaultID = attendanceDescID.Default.(func() uuid.UUID)
	classMixin := schema.Class{}.Mixin()
	classMixinFields0 := classMixin[0].Fields()
	_ = classMixinFields0
	classFields := schema.Class{}.Fields()
	_ = classFields
	// classDescCreatedAt is the schema descriptor for created_at field.
	classDescCreatedAt := classMixinFields0[0].Descriptor()
	// class.DefaultCreatedAt holds the default value on creation for the created_at field.
	class.DefaultCreatedAt = classDescCreatedAt.Default.(func() time.Time)
	// classDescUpdatedAt is the schema descriptor for updated_at field.
	classDescUpdatedAt := classMixinFields0[1].Descriptor()
	// class.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	class.DefaultUpdatedAt = classDescUpdatedAt.Default.(func() time.Time)
	// class.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	class.UpdateDefaultUpdatedAt = classDescUpdatedAt.UpdateDefault.(func() time.Time)
	// classDescName is the schema descriptor for name field.
	classDescName := classFields[1].Descriptor()
	// class.NameValidator is a validator for the "name" field. It is called by the builders before save.
	class.NameValidator = classDescName.Validators[0].(func(string) error)
	// classDescActive is the schema descriptor for active field.
	classDescActive := classFields[2].Descriptor()
	// class.DefaultActive holds the default value on creation for the active field.
	class.DefaultActive = classDescActive.Default.(bool)
	// classDescID is the schema descriptor for id field.
	classDescID := classFields[0].Descriptor()
	// class.DefaultID holds the default value on creation for the id field.
	class.DefaultID = classDescID.Default.(func() uuid.UUID)
	coursegradeMixin := schema.CourseGrade{}.Mixin()
	coursegradeMixinFields0 := coursegradeMixin[0].Fields()
	_ = coursegradeMixinFields0
	coursegradeFields := schema.CourseGrade{}.Fields()
	_ = coursegradeFields
	// coursegradeDescCreatedAt is the schema descriptor for created_at field.
	coursegradeDescCreatedAt := coursegradeMixinFields0[0].Descriptor()
	// coursegrade.DefaultCreatedAt holds the default value on creation for the created_at field.
	coursegrade.DefaultCreatedAt = coursegradeDescCreatedAt.Default.(func() time.Time)
	// coursegradeDescUpdatedAt is the schema descriptor for updated_at field.
	coursegradeDescUpdatedAt := coursegradeMixinFields0[1].Descriptor()
	// coursegrade.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	coursegrade.DefaultUpdatedAt = coursegradeDescUpdatedAt.Default.(func() time.Time)
	// coursegrade.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	coursegrade.UpdateDefaultUpdatedAt = coursegradeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// coursegradeDescActivityFirst is the schema descriptor for activity_first field.
	coursegradeDescActivityFirst := coursegradeFields[2].Descriptor()
	// coursegrade.ActivityFirstValidator is a validator for the "activity_first" field. It is called by the builders before save.
	coursegrade.ActivityFirstValidator = coursegradeDescActivityFirst.Validators[0].(func(int) error)
	// coursegradeDescActivitySecond is the schema descriptor for activity_second field.
	coursegradeDescActivitySecond := coursegradeFields[3].Descriptor()
	// coursegrade.ActivitySecondValidator is a validator for the "activity_second" field. It is called by the builders before save.
	coursegrade.ActivitySecondValidator = coursegradeDescActivitySecond.Validators[0].(func(int) error)
	// coursegradeDescWrittenFirst is the schema descriptor for written_first field.
	coursegradeDescWrittenFirst := coursegradeFields[4].Descriptor()
	// coursegrade.WrittenFirstValidator is a validator for the "written_first" field. It is called by the builders before save.
	coursegrade.WrittenFirstValidator = coursegradeDescWrittenFirst.Validators[0].(func(int) error)
	// coursegradeDescWrittenSecond is the schema descriptor for written_second field.
	coursegradeDescWrittenSecond := coursegradeFields[5].Descriptor()
	// coursegrade.WrittenSecondValidator is a validator for the "written_second" field. It is called by the builders before save.
	coursegrade.WrittenSecondValidator = coursegradeDescWrittenSecond.Validators[0].(func(int) error)
	// coursegradeDescCourseFinal is the schema descriptor for course_final field.
	coursegradeDescCourseFinal := coursegradeFields[6].Descriptor()
	// coursegrade.CourseFinalValidator is a validator for the "course_final" field. It is called by the builders before save.
	coursegrade.CourseFinalValidator = coursegradeDescCourseFinal.Validators[0].(func(int) error)
	// coursegradeDescYear is the schema descriptor for year field.
	coursegradeDescYear := coursegradeFields[7].Descriptor()
	// coursegrade.YearValidator is a validator for the "year" field. It is called by the builders before save.
	coursegrade.YearValidator = coursegradeDescYear.Validators[0].(func(string) error)
	// coursegradeDescID is the schema descriptor for id field.
	coursegradeDescID := coursegradeFields[0].Descriptor()
	// coursegrade.DefaultID holds the default value on creation for the id field.
	coursegrade.DefaultID = coursegradeDescID.Default.(func() uuid.UUID)
	gradeMixin := schema.Grade{}.Mixin()
	gradeMixinFields0 := gradeMixin[0].Fields()
	_ = gradeMixinFields0
	gradeFields := schema.Grade{}.Fields()
	_ = gradeFields
	// gradeDescCreatedAt is the schema descriptor for created_at field.
	gradeDescCreatedAt := gradeMixinFields0[0].Descriptor()
	// grade.DefaultCreatedAt holds the default value on creation for the created_at field.
	grade.DefaultCreatedAt = gradeDescCreatedAt.Default.(func() time.Time)
	// gradeDescUpdatedAt is the schema descriptor for updated_at field.
	gradeDescUpdatedAt := gradeMixinFields0[1].Descriptor()
	// grade.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	grade.DefaultUpdatedAt = gradeDescUpdatedAt.Default.(func() time.Time)
	// grade.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	grade.UpdateDefaultUpdatedAt = gradeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// gradeDescExamGrade is the schema descriptor for exam_grade field.
	gradeDescExamGrade := gradeFields[1].Descriptor()
	// grade.ExamGradeValidator is a validator for the "exam_grade" field. It is called by the builders before save.
	grade.ExamGradeValidator = gradeDescExamGrade.Validators[0].(func(int) error)
	// gradeDescID is the schema descriptor for id field.
	gradeDescID := gradeFields[0].Descriptor()
	// grade.DefaultID holds the default value on creation for the id field.
	grade.DefaultID = gradeDescID.Default.(func() uuid.UUID)
	groupMixin := schema.Group{}.Mixin()
	groupMixinFields0 := groupMixin[0].Fields()
	_ = groupMixinFields0
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescCreatedAt is the schema descriptor for created_at field.
	groupDescCreatedAt := groupMixinFields0[0].Descriptor()
	// group.DefaultCreatedAt holds the default value on creation for the created_at field.
	group.DefaultCreatedAt = groupDescCreatedAt.Default.(func() time.Time)
	// groupDescUpdatedAt is the schema descriptor for updated_at field.
	groupDescUpdatedAt := groupMixinFields0[1].Descriptor()
	// group.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	group.DefaultUpdatedAt = groupDescUpdatedAt.Default.(func() time.Time)
	// group.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	group.UpdateDefaultUpdatedAt = groupDescUpdatedAt.UpdateDefault.(func() time.Time)
	// groupDescActive is the schema descriptor for active field.
	groupDescActive := groupFields[3].Descriptor()
	// group.DefaultActive holds the default value on creation for the active field.
	group.DefaultActive = groupDescActive.Default.(bool)
	// groupDescID is the schema descriptor for id field.
	groupDescID := groupFields[0].Descriptor()
	// group.DefaultID holds the default value on creation for the id field.
	group.DefaultID = groupDescID.Default.(func() uuid.UUID)
	messageMixin := schema.Message{}.Mixin()
	messageMixinFields0 := messageMixin[0].Fields()
	_ = messageMixinFields0
	messageFields := schema.Message{}.Fields()
	_ = messageFields
	// messageDescCreatedAt is the schema descriptor for created_at field.
	messageDescCreatedAt := messageMixinFields0[0].Descriptor()
	// message.DefaultCreatedAt holds the default value on creation for the created_at field.
	message.DefaultCreatedAt = messageDescCreatedAt.Default.(func() time.Time)
	// messageDescUpdatedAt is the schema descriptor for updated_at field.
	messageDescUpdatedAt := messageMixinFields0[1].Descriptor()
	// message.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	message.DefaultUpdatedAt = messageDescUpdatedAt.Default.(func() time.Time)
	// message.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	message.UpdateDefaultUpdatedAt = messageDescUpdatedAt.UpdateDefault.(func() time.Time)
	// messageDescID is the schema descriptor for id field.
	messageDescID := messageFields[0].Descriptor()
	// message.DefaultID holds the default value on creation for the id field.
	message.DefaultID = messageDescID.Default.(func() uuid.UUID)
	notificationMixin := schema.Notification{}.Mixin()
	notificationMixinFields0 := notificationMixin[0].Fields()
	_ = notificationMixinFields0
	notificationFields := schema.Notification{}.Fields()
	_ = notificationFields
	// notificationDescCreatedAt is the schema descriptor for created_at field.
	notificationDescCreatedAt := notificationMixinFields0[0].Descriptor()
	// notification.DefaultCreatedAt holds the default value on creation for the created_at field.
	notification.DefaultCreatedAt = notificationDescCreatedAt.Default.(func() time.Time)
	// notificationDescUpdatedAt is the schema descriptor for updated_at field.
	notificationDescUpdatedAt := notificationMixinFields0[1].Descriptor()
	// notification.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	notification.DefaultUpdatedAt = notificationDescUpdatedAt.Default.(func() time.Time)
	// notification.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	notification.UpdateDefaultUpdatedAt = notificationDescUpdatedAt.UpdateDefault.(func() time.Time)
	// notificationDescRoute is the schema descriptor for route field.
	notificationDescRoute := notificationFields[4].Descriptor()
	// notification.RouteValidator is a validator for the "route" field. It is called by the builders before save.
	notification.RouteValidator = notificationDescRoute.Validators[0].(func(string) error)
	// notificationDescID is the schema descriptor for id field.
	notificationDescID := notificationFields[0].Descriptor()
	// notification.DefaultID holds the default value on creation for the id field.
	notification.DefaultID = notificationDescID.Default.(func() uuid.UUID)
	scheduleFields := schema.Schedule{}.Fields()
	_ = scheduleFields
	// scheduleDescWeekday is the schema descriptor for weekday field.
	scheduleDescWeekday := scheduleFields[1].Descriptor()
	// schedule.WeekdayValidator is a validator for the "weekday" field. It is called by the builders before save.
	schedule.WeekdayValidator = func() func(int) error {
		validators := scheduleDescWeekday.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
		}
		return func(weekday int) error {
			for _, fn := range fns {
				if err := fn(weekday); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// scheduleDescID is the schema descriptor for id field.
	scheduleDescID := scheduleFields[0].Descriptor()
	// schedule.DefaultID holds the default value on creation for the id field.
	schedule.DefaultID = scheduleDescID.Default.(func() uuid.UUID)
	schoolMixin := schema.School{}.Mixin()
	schoolMixinFields0 := schoolMixin[0].Fields()
	_ = schoolMixinFields0
	schoolFields := schema.School{}.Fields()
	_ = schoolFields
	// schoolDescCreatedAt is the schema descriptor for created_at field.
	schoolDescCreatedAt := schoolMixinFields0[0].Descriptor()
	// school.DefaultCreatedAt holds the default value on creation for the created_at field.
	school.DefaultCreatedAt = schoolDescCreatedAt.Default.(func() time.Time)
	// schoolDescUpdatedAt is the schema descriptor for updated_at field.
	schoolDescUpdatedAt := schoolMixinFields0[1].Descriptor()
	// school.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	school.DefaultUpdatedAt = schoolDescUpdatedAt.Default.(func() time.Time)
	// school.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	school.UpdateDefaultUpdatedAt = schoolDescUpdatedAt.UpdateDefault.(func() time.Time)
	// schoolDescName is the schema descriptor for name field.
	schoolDescName := schoolFields[1].Descriptor()
	// school.NameValidator is a validator for the "name" field. It is called by the builders before save.
	school.NameValidator = schoolDescName.Validators[0].(func(string) error)
	// schoolDescImage is the schema descriptor for image field.
	schoolDescImage := schoolFields[2].Descriptor()
	// school.ImageValidator is a validator for the "image" field. It is called by the builders before save.
	school.ImageValidator = schoolDescImage.Validators[0].(func(string) error)
	// schoolDescDirectory is the schema descriptor for directory field.
	schoolDescDirectory := schoolFields[3].Descriptor()
	// school.DirectoryValidator is a validator for the "directory" field. It is called by the builders before save.
	school.DirectoryValidator = schoolDescDirectory.Validators[0].(func(string) error)
	// schoolDescActive is the schema descriptor for active field.
	schoolDescActive := schoolFields[4].Descriptor()
	// school.DefaultActive holds the default value on creation for the active field.
	school.DefaultActive = schoolDescActive.Default.(bool)
	// schoolDescID is the schema descriptor for id field.
	schoolDescID := schoolFields[0].Descriptor()
	// school.DefaultID holds the default value on creation for the id field.
	school.DefaultID = schoolDescID.Default.(func() uuid.UUID)
	stageMixin := schema.Stage{}.Mixin()
	stageMixinFields0 := stageMixin[0].Fields()
	_ = stageMixinFields0
	stageFields := schema.Stage{}.Fields()
	_ = stageFields
	// stageDescCreatedAt is the schema descriptor for created_at field.
	stageDescCreatedAt := stageMixinFields0[0].Descriptor()
	// stage.DefaultCreatedAt holds the default value on creation for the created_at field.
	stage.DefaultCreatedAt = stageDescCreatedAt.Default.(func() time.Time)
	// stageDescUpdatedAt is the schema descriptor for updated_at field.
	stageDescUpdatedAt := stageMixinFields0[1].Descriptor()
	// stage.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	stage.DefaultUpdatedAt = stageDescUpdatedAt.Default.(func() time.Time)
	// stage.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	stage.UpdateDefaultUpdatedAt = stageDescUpdatedAt.UpdateDefault.(func() time.Time)
	// stageDescName is the schema descriptor for name field.
	stageDescName := stageFields[1].Descriptor()
	// stage.NameValidator is a validator for the "name" field. It is called by the builders before save.
	stage.NameValidator = stageDescName.Validators[0].(func(string) error)
	// stageDescTuitionAmount is the schema descriptor for tuition_amount field.
	stageDescTuitionAmount := stageFields[2].Descriptor()
	// stage.TuitionAmountValidator is a validator for the "tuition_amount" field. It is called by the builders before save.
	stage.TuitionAmountValidator = stageDescTuitionAmount.Validators[0].(func(int) error)
	// stageDescActive is the schema descriptor for active field.
	stageDescActive := stageFields[4].Descriptor()
	// stage.DefaultActive holds the default value on creation for the active field.
	stage.DefaultActive = stageDescActive.Default.(bool)
	// stageDescID is the schema descriptor for id field.
	stageDescID := stageFields[0].Descriptor()
	// stage.DefaultID holds the default value on creation for the id field.
	stage.DefaultID = stageDescID.Default.(func() uuid.UUID)
	tuitionpaymentMixin := schema.TuitionPayment{}.Mixin()
	tuitionpaymentMixinFields0 := tuitionpaymentMixin[0].Fields()
	_ = tuitionpaymentMixinFields0
	tuitionpaymentFields := schema.TuitionPayment{}.Fields()
	_ = tuitionpaymentFields
	// tuitionpaymentDescCreatedAt is the schema descriptor for created_at field.
	tuitionpaymentDescCreatedAt := tuitionpaymentMixinFields0[0].Descriptor()
	// tuitionpayment.DefaultCreatedAt holds the default value on creation for the created_at field.
	tuitionpayment.DefaultCreatedAt = tuitionpaymentDescCreatedAt.Default.(func() time.Time)
	// tuitionpaymentDescUpdatedAt is the schema descriptor for updated_at field.
	tuitionpaymentDescUpdatedAt := tuitionpaymentMixinFields0[1].Descriptor()
	// tuitionpayment.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	tuitionpayment.DefaultUpdatedAt = tuitionpaymentDescUpdatedAt.Default.(func() time.Time)
	// tuitionpayment.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	tuitionpayment.UpdateDefaultUpdatedAt = tuitionpaymentDescUpdatedAt.UpdateDefault.(func() time.Time)
	// tuitionpaymentDescYear is the schema descriptor for year field.
	tuitionpaymentDescYear := tuitionpaymentFields[1].Descriptor()
	// tuitionpayment.YearValidator is a validator for the "year" field. It is called by the builders before save.
	tuitionpayment.YearValidator = tuitionpaymentDescYear.Validators[0].(func(string) error)
	// tuitionpaymentDescID is the schema descriptor for id field.
	tuitionpaymentDescID := tuitionpaymentFields[0].Descriptor()
	// tuitionpayment.DefaultID holds the default value on creation for the id field.
	tuitionpayment.DefaultID = tuitionpaymentDescID.Default.(func() uuid.UUID)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[2].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[3].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescPhone is the schema descriptor for phone field.
	userDescPhone := userFields[4].Descriptor()
	// user.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	user.PhoneValidator = userDescPhone.Validators[0].(func(string) error)
	// userDescDirectory is the schema descriptor for directory field.
	userDescDirectory := userFields[6].Descriptor()
	// user.DirectoryValidator is a validator for the "directory" field. It is called by the builders before save.
	user.DirectoryValidator = userDescDirectory.Validators[0].(func(string) error)
	// userDescTokenVersion is the schema descriptor for token_version field.
	userDescTokenVersion := userFields[7].Descriptor()
	// user.DefaultTokenVersion holds the default value on creation for the token_version field.
	user.DefaultTokenVersion = userDescTokenVersion.Default.(int)
	// userDescActive is the schema descriptor for active field.
	userDescActive := userFields[10].Descriptor()
	// user.DefaultActive holds the default value on creation for the active field.
	user.DefaultActive = userDescActive.Default.(bool)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
