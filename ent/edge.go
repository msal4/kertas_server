// Code generated by entc, DO NOT EDIT.

package ent

import "context"

func (a *Assignment) Class(ctx context.Context) (*Class, error) {
	result, err := a.Edges.ClassOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryClass().Only(ctx)
	}
	return result, err
}

func (a *Assignment) Submissions(ctx context.Context) ([]*AssignmentSubmission, error) {
	result, err := a.Edges.SubmissionsOrErr()
	if IsNotLoaded(err) {
		result, err = a.QuerySubmissions().All(ctx)
	}
	return result, err
}

func (a *Assignment) Grades(ctx context.Context) ([]*Grade, error) {
	result, err := a.Edges.GradesOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryGrades().All(ctx)
	}
	return result, err
}

func (as *AssignmentSubmission) Student(ctx context.Context) (*User, error) {
	result, err := as.Edges.StudentOrErr()
	if IsNotLoaded(err) {
		result, err = as.QueryStudent().Only(ctx)
	}
	return result, err
}

func (as *AssignmentSubmission) Assignment(ctx context.Context) (*Assignment, error) {
	result, err := as.Edges.AssignmentOrErr()
	if IsNotLoaded(err) {
		result, err = as.QueryAssignment().Only(ctx)
	}
	return result, err
}

func (a *Attendance) Class(ctx context.Context) (*Class, error) {
	result, err := a.Edges.ClassOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryClass().Only(ctx)
	}
	return result, err
}

func (a *Attendance) Student(ctx context.Context) (*User, error) {
	result, err := a.Edges.StudentOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryStudent().Only(ctx)
	}
	return result, err
}

func (c *Class) Stage(ctx context.Context) (*Stage, error) {
	result, err := c.Edges.StageOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryStage().Only(ctx)
	}
	return result, err
}

func (c *Class) Teacher(ctx context.Context) (*User, error) {
	result, err := c.Edges.TeacherOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryTeacher().Only(ctx)
	}
	return result, err
}

func (c *Class) Group(ctx context.Context) (*Group, error) {
	result, err := c.Edges.GroupOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryGroup().Only(ctx)
	}
	return result, err
}

func (c *Class) Assignments(ctx context.Context) ([]*Assignment, error) {
	result, err := c.Edges.AssignmentsOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryAssignments().All(ctx)
	}
	return result, err
}

func (c *Class) Attendances(ctx context.Context) ([]*Attendance, error) {
	result, err := c.Edges.AttendancesOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryAttendances().All(ctx)
	}
	return result, err
}

func (c *Class) Schedules(ctx context.Context) ([]*Schedule, error) {
	result, err := c.Edges.SchedulesOrErr()
	if IsNotLoaded(err) {
		result, err = c.QuerySchedules().All(ctx)
	}
	return result, err
}

func (gr *Grade) Student(ctx context.Context) (*User, error) {
	result, err := gr.Edges.StudentOrErr()
	if IsNotLoaded(err) {
		result, err = gr.QueryStudent().Only(ctx)
	}
	return result, err
}

func (gr *Grade) Exam(ctx context.Context) (*Assignment, error) {
	result, err := gr.Edges.ExamOrErr()
	if IsNotLoaded(err) {
		result, err = gr.QueryExam().Only(ctx)
	}
	return result, err
}

func (gr *Group) Class(ctx context.Context) (*Class, error) {
	result, err := gr.Edges.ClassOrErr()
	if IsNotLoaded(err) {
		result, err = gr.QueryClass().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (gr *Group) Messages(ctx context.Context) ([]*Message, error) {
	result, err := gr.Edges.MessagesOrErr()
	if IsNotLoaded(err) {
		result, err = gr.QueryMessages().All(ctx)
	}
	return result, err
}

func (m *Message) Group(ctx context.Context) (*Group, error) {
	result, err := m.Edges.GroupOrErr()
	if IsNotLoaded(err) {
		result, err = m.QueryGroup().Only(ctx)
	}
	return result, err
}

func (m *Message) Owner(ctx context.Context) (*User, error) {
	result, err := m.Edges.OwnerOrErr()
	if IsNotLoaded(err) {
		result, err = m.QueryOwner().Only(ctx)
	}
	return result, err
}

func (s *Schedule) Class(ctx context.Context) (*Class, error) {
	result, err := s.Edges.ClassOrErr()
	if IsNotLoaded(err) {
		result, err = s.QueryClass().Only(ctx)
	}
	return result, err
}

func (s *School) Users(ctx context.Context) ([]*User, error) {
	result, err := s.Edges.UsersOrErr()
	if IsNotLoaded(err) {
		result, err = s.QueryUsers().All(ctx)
	}
	return result, err
}

func (s *School) Stages(ctx context.Context) ([]*Stage, error) {
	result, err := s.Edges.StagesOrErr()
	if IsNotLoaded(err) {
		result, err = s.QueryStages().All(ctx)
	}
	return result, err
}

func (s *Stage) School(ctx context.Context) (*School, error) {
	result, err := s.Edges.SchoolOrErr()
	if IsNotLoaded(err) {
		result, err = s.QuerySchool().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (s *Stage) Classes(ctx context.Context) ([]*Class, error) {
	result, err := s.Edges.ClassesOrErr()
	if IsNotLoaded(err) {
		result, err = s.QueryClasses().All(ctx)
	}
	return result, err
}

func (s *Stage) Payments(ctx context.Context) ([]*TuitionPayment, error) {
	result, err := s.Edges.PaymentsOrErr()
	if IsNotLoaded(err) {
		result, err = s.QueryPayments().All(ctx)
	}
	return result, err
}

func (s *Stage) Students(ctx context.Context) ([]*User, error) {
	result, err := s.Edges.StudentsOrErr()
	if IsNotLoaded(err) {
		result, err = s.QueryStudents().All(ctx)
	}
	return result, err
}

func (tp *TuitionPayment) Student(ctx context.Context) (*User, error) {
	result, err := tp.Edges.StudentOrErr()
	if IsNotLoaded(err) {
		result, err = tp.QueryStudent().Only(ctx)
	}
	return result, err
}

func (tp *TuitionPayment) Stage(ctx context.Context) (*Stage, error) {
	result, err := tp.Edges.StageOrErr()
	if IsNotLoaded(err) {
		result, err = tp.QueryStage().Only(ctx)
	}
	return result, err
}

func (u *User) Stage(ctx context.Context) (*Stage, error) {
	result, err := u.Edges.StageOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryStage().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) School(ctx context.Context) (*School, error) {
	result, err := u.Edges.SchoolOrErr()
	if IsNotLoaded(err) {
		result, err = u.QuerySchool().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) Classes(ctx context.Context) ([]*Class, error) {
	result, err := u.Edges.ClassesOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryClasses().All(ctx)
	}
	return result, err
}

func (u *User) Messages(ctx context.Context) ([]*Message, error) {
	result, err := u.Edges.MessagesOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryMessages().All(ctx)
	}
	return result, err
}

func (u *User) Submissions(ctx context.Context) ([]*AssignmentSubmission, error) {
	result, err := u.Edges.SubmissionsOrErr()
	if IsNotLoaded(err) {
		result, err = u.QuerySubmissions().All(ctx)
	}
	return result, err
}

func (u *User) Attendances(ctx context.Context) ([]*Attendance, error) {
	result, err := u.Edges.AttendancesOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryAttendances().All(ctx)
	}
	return result, err
}

func (u *User) Payments(ctx context.Context) ([]*TuitionPayment, error) {
	result, err := u.Edges.PaymentsOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryPayments().All(ctx)
	}
	return result, err
}

func (u *User) Grades(ctx context.Context) ([]*Grade, error) {
	result, err := u.Edges.GradesOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryGrades().All(ctx)
	}
	return result, err
}