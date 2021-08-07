// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/msal4/hassah_school_server/ent/assignment"
	"github.com/msal4/hassah_school_server/ent/assignmentsubmission"
	"github.com/msal4/hassah_school_server/ent/attendance"
	"github.com/msal4/hassah_school_server/ent/class"
	"github.com/msal4/hassah_school_server/ent/grade"
	"github.com/msal4/hassah_school_server/ent/group"
	"github.com/msal4/hassah_school_server/ent/message"
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
	// assignmentDescCreateTime is the schema descriptor for create_time field.
	assignmentDescCreateTime := assignmentMixinFields0[0].Descriptor()
	// assignment.DefaultCreateTime holds the default value on creation for the create_time field.
	assignment.DefaultCreateTime = assignmentDescCreateTime.Default.(func() time.Time)
	// assignmentDescUpdateTime is the schema descriptor for update_time field.
	assignmentDescUpdateTime := assignmentMixinFields0[1].Descriptor()
	// assignment.DefaultUpdateTime holds the default value on creation for the update_time field.
	assignment.DefaultUpdateTime = assignmentDescUpdateTime.Default.(func() time.Time)
	// assignment.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	assignment.UpdateDefaultUpdateTime = assignmentDescUpdateTime.UpdateDefault.(func() time.Time)
	// assignmentDescName is the schema descriptor for name field.
	assignmentDescName := assignmentFields[0].Descriptor()
	// assignment.NameValidator is a validator for the "name" field. It is called by the builders before save.
	assignment.NameValidator = assignmentDescName.Validators[0].(func(string) error)
	// assignmentDescIsExam is the schema descriptor for is_exam field.
	assignmentDescIsExam := assignmentFields[2].Descriptor()
	// assignment.DefaultIsExam holds the default value on creation for the is_exam field.
	assignment.DefaultIsExam = assignmentDescIsExam.Default.(bool)
	assignmentsubmissionMixin := schema.AssignmentSubmission{}.Mixin()
	assignmentsubmissionMixinFields0 := assignmentsubmissionMixin[0].Fields()
	_ = assignmentsubmissionMixinFields0
	assignmentsubmissionFields := schema.AssignmentSubmission{}.Fields()
	_ = assignmentsubmissionFields
	// assignmentsubmissionDescCreateTime is the schema descriptor for create_time field.
	assignmentsubmissionDescCreateTime := assignmentsubmissionMixinFields0[0].Descriptor()
	// assignmentsubmission.DefaultCreateTime holds the default value on creation for the create_time field.
	assignmentsubmission.DefaultCreateTime = assignmentsubmissionDescCreateTime.Default.(func() time.Time)
	// assignmentsubmissionDescUpdateTime is the schema descriptor for update_time field.
	assignmentsubmissionDescUpdateTime := assignmentsubmissionMixinFields0[1].Descriptor()
	// assignmentsubmission.DefaultUpdateTime holds the default value on creation for the update_time field.
	assignmentsubmission.DefaultUpdateTime = assignmentsubmissionDescUpdateTime.Default.(func() time.Time)
	// assignmentsubmission.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	assignmentsubmission.UpdateDefaultUpdateTime = assignmentsubmissionDescUpdateTime.UpdateDefault.(func() time.Time)
	attendanceMixin := schema.Attendance{}.Mixin()
	attendanceMixinFields0 := attendanceMixin[0].Fields()
	_ = attendanceMixinFields0
	attendanceFields := schema.Attendance{}.Fields()
	_ = attendanceFields
	// attendanceDescCreateTime is the schema descriptor for create_time field.
	attendanceDescCreateTime := attendanceMixinFields0[0].Descriptor()
	// attendance.DefaultCreateTime holds the default value on creation for the create_time field.
	attendance.DefaultCreateTime = attendanceDescCreateTime.Default.(func() time.Time)
	// attendanceDescUpdateTime is the schema descriptor for update_time field.
	attendanceDescUpdateTime := attendanceMixinFields0[1].Descriptor()
	// attendance.DefaultUpdateTime holds the default value on creation for the update_time field.
	attendance.DefaultUpdateTime = attendanceDescUpdateTime.Default.(func() time.Time)
	// attendance.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	attendance.UpdateDefaultUpdateTime = attendanceDescUpdateTime.UpdateDefault.(func() time.Time)
	classMixin := schema.Class{}.Mixin()
	classMixinFields0 := classMixin[0].Fields()
	_ = classMixinFields0
	classFields := schema.Class{}.Fields()
	_ = classFields
	// classDescCreateTime is the schema descriptor for create_time field.
	classDescCreateTime := classMixinFields0[0].Descriptor()
	// class.DefaultCreateTime holds the default value on creation for the create_time field.
	class.DefaultCreateTime = classDescCreateTime.Default.(func() time.Time)
	// classDescUpdateTime is the schema descriptor for update_time field.
	classDescUpdateTime := classMixinFields0[1].Descriptor()
	// class.DefaultUpdateTime holds the default value on creation for the update_time field.
	class.DefaultUpdateTime = classDescUpdateTime.Default.(func() time.Time)
	// class.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	class.UpdateDefaultUpdateTime = classDescUpdateTime.UpdateDefault.(func() time.Time)
	// classDescName is the schema descriptor for name field.
	classDescName := classFields[0].Descriptor()
	// class.NameValidator is a validator for the "name" field. It is called by the builders before save.
	class.NameValidator = classDescName.Validators[0].(func(string) error)
	gradeMixin := schema.Grade{}.Mixin()
	gradeMixinFields0 := gradeMixin[0].Fields()
	_ = gradeMixinFields0
	gradeFields := schema.Grade{}.Fields()
	_ = gradeFields
	// gradeDescCreateTime is the schema descriptor for create_time field.
	gradeDescCreateTime := gradeMixinFields0[0].Descriptor()
	// grade.DefaultCreateTime holds the default value on creation for the create_time field.
	grade.DefaultCreateTime = gradeDescCreateTime.Default.(func() time.Time)
	// gradeDescUpdateTime is the schema descriptor for update_time field.
	gradeDescUpdateTime := gradeMixinFields0[1].Descriptor()
	// grade.DefaultUpdateTime holds the default value on creation for the update_time field.
	grade.DefaultUpdateTime = gradeDescUpdateTime.Default.(func() time.Time)
	// grade.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	grade.UpdateDefaultUpdateTime = gradeDescUpdateTime.UpdateDefault.(func() time.Time)
	// gradeDescExamGrade is the schema descriptor for exam_grade field.
	gradeDescExamGrade := gradeFields[0].Descriptor()
	// grade.ExamGradeValidator is a validator for the "exam_grade" field. It is called by the builders before save.
	grade.ExamGradeValidator = gradeDescExamGrade.Validators[0].(func(float64) error)
	groupMixin := schema.Group{}.Mixin()
	groupMixinFields0 := groupMixin[0].Fields()
	_ = groupMixinFields0
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescCreateTime is the schema descriptor for create_time field.
	groupDescCreateTime := groupMixinFields0[0].Descriptor()
	// group.DefaultCreateTime holds the default value on creation for the create_time field.
	group.DefaultCreateTime = groupDescCreateTime.Default.(func() time.Time)
	// groupDescUpdateTime is the schema descriptor for update_time field.
	groupDescUpdateTime := groupMixinFields0[1].Descriptor()
	// group.DefaultUpdateTime holds the default value on creation for the update_time field.
	group.DefaultUpdateTime = groupDescUpdateTime.Default.(func() time.Time)
	// group.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	group.UpdateDefaultUpdateTime = groupDescUpdateTime.UpdateDefault.(func() time.Time)
	messageMixin := schema.Message{}.Mixin()
	messageMixinFields0 := messageMixin[0].Fields()
	_ = messageMixinFields0
	messageFields := schema.Message{}.Fields()
	_ = messageFields
	// messageDescCreateTime is the schema descriptor for create_time field.
	messageDescCreateTime := messageMixinFields0[0].Descriptor()
	// message.DefaultCreateTime holds the default value on creation for the create_time field.
	message.DefaultCreateTime = messageDescCreateTime.Default.(func() time.Time)
	// messageDescUpdateTime is the schema descriptor for update_time field.
	messageDescUpdateTime := messageMixinFields0[1].Descriptor()
	// message.DefaultUpdateTime holds the default value on creation for the update_time field.
	message.DefaultUpdateTime = messageDescUpdateTime.Default.(func() time.Time)
	// message.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	message.UpdateDefaultUpdateTime = messageDescUpdateTime.UpdateDefault.(func() time.Time)
	scheduleFields := schema.Schedule{}.Fields()
	_ = scheduleFields
	// scheduleDescWeekday is the schema descriptor for weekday field.
	scheduleDescWeekday := scheduleFields[0].Descriptor()
	// schedule.WeekdayValidator is a validator for the "weekday" field. It is called by the builders before save.
	schedule.WeekdayValidator = scheduleDescWeekday.Validators[0].(func(uint8) error)
	// scheduleDescDuration is the schema descriptor for duration field.
	scheduleDescDuration := scheduleFields[2].Descriptor()
	// schedule.DefaultDuration holds the default value on creation for the duration field.
	schedule.DefaultDuration = scheduleDescDuration.Default.(int)
	schoolMixin := schema.School{}.Mixin()
	schoolMixinFields0 := schoolMixin[0].Fields()
	_ = schoolMixinFields0
	schoolFields := schema.School{}.Fields()
	_ = schoolFields
	// schoolDescCreateTime is the schema descriptor for create_time field.
	schoolDescCreateTime := schoolMixinFields0[0].Descriptor()
	// school.DefaultCreateTime holds the default value on creation for the create_time field.
	school.DefaultCreateTime = schoolDescCreateTime.Default.(func() time.Time)
	// schoolDescUpdateTime is the schema descriptor for update_time field.
	schoolDescUpdateTime := schoolMixinFields0[1].Descriptor()
	// school.DefaultUpdateTime holds the default value on creation for the update_time field.
	school.DefaultUpdateTime = schoolDescUpdateTime.Default.(func() time.Time)
	// school.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	school.UpdateDefaultUpdateTime = schoolDescUpdateTime.UpdateDefault.(func() time.Time)
	// schoolDescName is the schema descriptor for name field.
	schoolDescName := schoolFields[0].Descriptor()
	// school.NameValidator is a validator for the "name" field. It is called by the builders before save.
	school.NameValidator = schoolDescName.Validators[0].(func(string) error)
	// schoolDescImage is the schema descriptor for image field.
	schoolDescImage := schoolFields[1].Descriptor()
	// school.ImageValidator is a validator for the "image" field. It is called by the builders before save.
	school.ImageValidator = schoolDescImage.Validators[0].(func(string) error)
	stageMixin := schema.Stage{}.Mixin()
	stageMixinFields0 := stageMixin[0].Fields()
	_ = stageMixinFields0
	stageFields := schema.Stage{}.Fields()
	_ = stageFields
	// stageDescCreateTime is the schema descriptor for create_time field.
	stageDescCreateTime := stageMixinFields0[0].Descriptor()
	// stage.DefaultCreateTime holds the default value on creation for the create_time field.
	stage.DefaultCreateTime = stageDescCreateTime.Default.(func() time.Time)
	// stageDescUpdateTime is the schema descriptor for update_time field.
	stageDescUpdateTime := stageMixinFields0[1].Descriptor()
	// stage.DefaultUpdateTime holds the default value on creation for the update_time field.
	stage.DefaultUpdateTime = stageDescUpdateTime.Default.(func() time.Time)
	// stage.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	stage.UpdateDefaultUpdateTime = stageDescUpdateTime.UpdateDefault.(func() time.Time)
	// stageDescName is the schema descriptor for name field.
	stageDescName := stageFields[0].Descriptor()
	// stage.NameValidator is a validator for the "name" field. It is called by the builders before save.
	stage.NameValidator = stageDescName.Validators[0].(func(string) error)
	tuitionpaymentMixin := schema.TuitionPayment{}.Mixin()
	tuitionpaymentMixinFields0 := tuitionpaymentMixin[0].Fields()
	_ = tuitionpaymentMixinFields0
	tuitionpaymentFields := schema.TuitionPayment{}.Fields()
	_ = tuitionpaymentFields
	// tuitionpaymentDescCreateTime is the schema descriptor for create_time field.
	tuitionpaymentDescCreateTime := tuitionpaymentMixinFields0[0].Descriptor()
	// tuitionpayment.DefaultCreateTime holds the default value on creation for the create_time field.
	tuitionpayment.DefaultCreateTime = tuitionpaymentDescCreateTime.Default.(func() time.Time)
	// tuitionpaymentDescUpdateTime is the schema descriptor for update_time field.
	tuitionpaymentDescUpdateTime := tuitionpaymentMixinFields0[1].Descriptor()
	// tuitionpayment.DefaultUpdateTime holds the default value on creation for the update_time field.
	tuitionpayment.DefaultUpdateTime = tuitionpaymentDescUpdateTime.Default.(func() time.Time)
	// tuitionpayment.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	tuitionpayment.UpdateDefaultUpdateTime = tuitionpaymentDescUpdateTime.UpdateDefault.(func() time.Time)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields0[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields0[1].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[1].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescPhone is the schema descriptor for phone field.
	userDescPhone := userFields[3].Descriptor()
	// user.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	user.PhoneValidator = userDescPhone.Validators[0].(func(string) error)
	// userDescTokenVersion is the schema descriptor for token_version field.
	userDescTokenVersion := userFields[5].Descriptor()
	// user.DefaultTokenVersion holds the default value on creation for the token_version field.
	user.DefaultTokenVersion = userDescTokenVersion.Default.(int)
}
