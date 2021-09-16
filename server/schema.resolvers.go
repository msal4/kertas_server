package server

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/auth"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/assignment"
	"github.com/msal4/hassah_school_server/ent/class"
	"github.com/msal4/hassah_school_server/ent/group"
	"github.com/msal4/hassah_school_server/ent/message"
	"github.com/msal4/hassah_school_server/ent/stage"
	"github.com/msal4/hassah_school_server/ent/user"
	"github.com/msal4/hassah_school_server/server/generated"
	"github.com/msal4/hassah_school_server/server/model"
	"github.com/msal4/hassah_school_server/service"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *assignmentResolver) Submissions(ctx context.Context, obj *ent.Assignment, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.AssignmentSubmissionOrder, where *ent.AssignmentSubmissionWhereInput) (*ent.AssignmentSubmissionConnection, error) {
	return obj.QuerySubmissions().Paginate(ctx, after, first, before, last,
		ent.WithAssignmentSubmissionOrder(orderBy), ent.WithAssignmentSubmissionFilter(where.Filter))
}

func (r *classResolver) Assignments(ctx context.Context, obj *ent.Class, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.AssignmentOrder, where *ent.AssignmentWhereInput) (*ent.AssignmentConnection, error) {
	return obj.QueryAssignments().Where(assignment.DeletedAtIsNil()).Paginate(ctx, after, first, before, last, ent.WithAssignmentOrder(orderBy), ent.WithAssignmentFilter(where.Filter))
}

func (r *classResolver) Attendances(ctx context.Context, obj *ent.Class, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.AttendanceOrder, where *ent.AttendanceWhereInput) (*ent.AttendanceConnection, error) {
	return obj.QueryAttendances().Paginate(ctx, after, first, before, last, ent.WithAttendanceOrder(orderBy), ent.WithAttendanceFilter(where.Filter))
}

func (r *classResolver) Schedules(ctx context.Context, obj *ent.Class, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ScheduleOrder, where *ent.ScheduleWhereInput) (*ent.ScheduleConnection, error) {
	return obj.QuerySchedules().Paginate(ctx, after, first, before, last, ent.WithScheduleOrder(orderBy), ent.WithScheduleFilter(where.Filter))
}

func (r *classResolver) CourseGrades(ctx context.Context, obj *ent.Class, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.CourseGradeOrder, where *ent.CourseGradeWhereInput) (*ent.CourseGradeConnection, error) {
	return obj.QueryCourseGrades().Paginate(ctx, after, first, before, last, ent.WithCourseGradeOrder(orderBy), ent.WithCourseGradeFilter(where.Filter))
}

func (r *groupResolver) Messages(ctx context.Context, obj *ent.Group, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.MessageOrder, where *ent.MessageWhereInput) (*ent.MessageConnection, error) {
	return obj.QueryMessages().Where(message.DeletedAtIsNil()).Paginate(ctx, after, first, before, last, ent.WithMessageOrder(orderBy), ent.WithMessageFilter(where.Filter))
}

func (r *mutationResolver) AddSchool(ctx context.Context, input model.AddSchoolInput) (*ent.School, error) {
	if !auth.IsSuperAdmin(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.AddSchool(ctx, input)
}

func (r *mutationResolver) UpdateSchool(ctx context.Context, id uuid.UUID, input model.UpdateSchoolInput) (*ent.School, error) {
	if !auth.IsSuperAdmin(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.UpdateSchool(ctx, id, input)
}

func (r *mutationResolver) DeleteSchool(ctx context.Context, id uuid.UUID) (bool, error) {
	if !auth.IsSuperAdmin(ctx) {
		return false, auth.UnauthorizedErr
	}

	return true, r.s.DeleteSchool(ctx, id)
}

func (r *mutationResolver) DeleteSchoolPermanently(ctx context.Context, id uuid.UUID) (bool, error) {
	if !auth.IsSuperAdmin(ctx) {
		return false, auth.UnauthorizedErr
	}

	return true, r.s.DeleteSchoolPermanently(ctx, id)
}

func (r *mutationResolver) AddUser(ctx context.Context, input model.AddUserInput) (*ent.User, error) {
	if !auth.IsAdmin(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.AddUser(ctx, input)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id uuid.UUID, input model.UpdateUserInput) (*ent.User, error) {
	if !auth.IsAdmin(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.UpdateUser(ctx, id, input)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id uuid.UUID) (bool, error) {
	if !auth.IsAdmin(ctx) {
		return false, auth.UnauthorizedErr
	}

	return true, r.s.DeleteUser(ctx, id)
}

func (r *mutationResolver) DeleteUserPermanently(ctx context.Context, id uuid.UUID) (bool, error) {
	if !auth.IsAdmin(ctx) {
		return false, auth.UnauthorizedErr
	}

	return true, r.s.DeleteUserPermanently(ctx, id)
}

func (r *mutationResolver) AddStage(ctx context.Context, input model.AddStageInput) (*ent.Stage, error) {
	if !auth.IsAdmin(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.AddStage(ctx, input)
}

func (r *mutationResolver) UpdateStage(ctx context.Context, id uuid.UUID, input model.UpdateStageInput) (*ent.Stage, error) {
	if !auth.IsAdmin(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.UpdateStage(ctx, id, input)
}

func (r *mutationResolver) DeleteStage(ctx context.Context, id uuid.UUID) (bool, error) {
	if !auth.IsAdmin(ctx) {
		return false, auth.UnauthorizedErr
	}

	return true, r.s.DeleteStage(ctx, id)
}

func (r *mutationResolver) DeleteStagePermanently(ctx context.Context, id uuid.UUID) (bool, error) {
	if !auth.IsAdmin(ctx) {
		return false, auth.UnauthorizedErr
	}

	return true, r.s.DeleteStagePermanently(ctx, id)
}

func getLoginErrs(ctx context.Context, err error) error {
	if err == nil {
		return nil
	}

	exts := map[string]interface{}{}

	switch err {
	case service.NotFoundErr:
		exts["code"] = "NOT_FOUND"
	case service.InvalidCredsErr:
		exts["code"] = "INVALID_CREDS"
	case service.NotAllowedErr:
		exts["code"] = "NOT_ALLOWED"
	}

	return &gqlerror.Error{
		Message:    err.Error(),
		Path:       graphql.GetPath(ctx),
		Extensions: exts,
	}
}

func (r *mutationResolver) LoginAdmin(ctx context.Context, input model.LoginInput) (*model.AuthData, error) {
	data, err := r.s.LoginAdmin(ctx, input)
	return data, getLoginErrs(ctx, err)
}

func (r *mutationResolver) LoginUser(ctx context.Context, input model.LoginInput) (*model.AuthData, error) {
	data, err := r.s.LoginUser(ctx, input)
	return data, getLoginErrs(ctx, err)
}

func (r *mutationResolver) RefreshTokens(ctx context.Context, token string) (*model.AuthData, error) {
	data, err := r.s.RefreshTokens(ctx, token)
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
			Extensions: map[string]interface{}{
				"code": "INVALID_TOKEN",
			},
		}
	}

	return data, nil
}

func (r *mutationResolver) PostMessage(ctx context.Context, input model.PostMessageInput) (*ent.Message, error) {
	u, ok := auth.UserForContext(ctx)
	if !ok {
		return nil, auth.UnauthorizedErr
	}
	return r.s.PostMessage(ctx, u.ID, input)
}

func (r *mutationResolver) AddGroup(ctx context.Context, input model.AddGroupInput) (*ent.Group, error) {
	u, ok := auth.UserForContext(ctx)
	if !ok {
		return nil, auth.UnauthorizedErr
	}

	return r.s.AddGroup(ctx, service.AddGroupInput{
		Name:    input.Name,
		Active:  input.Active,
		UserIDs: []uuid.UUID{u.ID, input.UserID},
	})
}

func (r *mutationResolver) UpdateGroup(ctx context.Context, id uuid.UUID, input model.UpdateGroupInput) (*ent.Group, error) {
	if !auth.IsAdmin(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.UpdateGroup(ctx, id, input)
}

func (r *mutationResolver) DeleteGroup(ctx context.Context, id uuid.UUID) (bool, error) {
	if !auth.IsAdmin(ctx) {
		return false, auth.UnauthorizedErr
	}

	return true, r.s.DeleteGroup(ctx, id)
}

func (r *mutationResolver) AddClass(ctx context.Context, input model.AddClassInput) (*ent.Class, error) {
	if !auth.IsAdmin(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.AddClass(ctx, input)
}

func (r *mutationResolver) UpdateClass(ctx context.Context, id uuid.UUID, input model.UpdateClassInput) (*ent.Class, error) {
	if !auth.IsAdmin(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.UpdateClass(ctx, id, input)
}

func (r *mutationResolver) DeleteClass(ctx context.Context, id uuid.UUID) (bool, error) {
	if !auth.IsAdmin(ctx) {
		return false, auth.UnauthorizedErr
	}

	return true, r.s.DeleteClass(ctx, id)
}

func (r *mutationResolver) AddAssignment(ctx context.Context, input model.AddAssignmentInput) (*ent.Assignment, error) {
	if !auth.IsAuthorized(ctx, user.RoleSuperAdmin, user.RoleSchoolAdmin, user.RoleTeacher) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.AddAssignment(ctx, input)
}

func (r *mutationResolver) UpdateAssignment(ctx context.Context, id uuid.UUID, input model.UpdateAssignmentInput) (*ent.Assignment, error) {
	if !auth.IsAuthorized(ctx, user.RoleSuperAdmin, user.RoleSchoolAdmin, user.RoleTeacher) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.UpdateAssignment(ctx, id, input)
}

func (r *mutationResolver) DeleteAssignment(ctx context.Context, id uuid.UUID) (bool, error) {
	if !auth.IsAuthorized(ctx, user.RoleSuperAdmin, user.RoleSchoolAdmin, user.RoleTeacher) {
		return false, auth.UnauthorizedErr
	}

	return true, r.s.DeleteAssignment(ctx, id)
}

func (r *mutationResolver) AddAssignmentSubmission(ctx context.Context, input model.AddAssignmentSubmissionInput) (*ent.AssignmentSubmission, error) {
	u, ok := auth.UserForContext(ctx)
	if !ok || u.Role != user.RoleStudent {
		return nil, auth.UnauthorizedErr
	}

	return r.s.AddAssignmentSubmission(ctx, u.ID, input)
}

func (r *mutationResolver) UpdateAssignmentSubmission(ctx context.Context, id uuid.UUID, input model.UpdateAssignmentSubmissionInput) (*ent.AssignmentSubmission, error) {
	if !auth.IsStudent(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.UpdateAssignmentSubmission(ctx, id, input)
}

func (r *mutationResolver) DeleteAssignmentSubmissionFile(ctx context.Context, id uuid.UUID, index int) (*ent.AssignmentSubmission, error) {
	if !auth.IsStudent(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.DeleteAssignmentSubmissionFile(ctx, id, index)
}

func (r *mutationResolver) DeleteAssignmentSubmission(ctx context.Context, id uuid.UUID) (bool, error) {
	if !auth.IsAuthorized(ctx) {
		return false, auth.UnauthorizedErr
	}

	return true, r.s.DeleteAssignmentSubmission(ctx, id)
}

func (r *mutationResolver) AddSchedule(ctx context.Context, input model.AddScheduleInput) (*ent.Schedule, error) {
	if !auth.IsAdmin(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.AddSchedule(ctx, input)
}

func (r *mutationResolver) UpdateSchedule(ctx context.Context, id uuid.UUID, input model.UpdateScheduleInput) (*ent.Schedule, error) {
	if !auth.IsAdmin(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.UpdateSchedule(ctx, id, input)
}

func (r *mutationResolver) DeleteSchedule(ctx context.Context, id uuid.UUID) (bool, error) {
	if !auth.IsAdmin(ctx) {
		return false, auth.UnauthorizedErr
	}

	return true, r.s.DeleteSchedule(ctx, id)
}

func (r *mutationResolver) AddCourseGrade(ctx context.Context, input model.AddCourseGradeInput) (*ent.CourseGrade, error) {
	if !auth.IsAuthorized(ctx, user.RoleSuperAdmin, user.RoleSchoolAdmin, user.RoleTeacher) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.AddCourseGrade(ctx, input)
}

func (r *mutationResolver) UpdateCourseGrade(ctx context.Context, id uuid.UUID, input model.UpdateCourseGradeInput) (*ent.CourseGrade, error) {
	if !auth.IsAuthorized(ctx, user.RoleSuperAdmin, user.RoleSchoolAdmin, user.RoleTeacher) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.UpdateCourseGrade(ctx, id, input)
}

func (r *mutationResolver) DeleteCourseGrade(ctx context.Context, id uuid.UUID) (bool, error) {
	if !auth.IsAuthorized(ctx, user.RoleSuperAdmin, user.RoleSchoolAdmin, user.RoleTeacher) {
		return false, auth.UnauthorizedErr
	}

	return true, r.s.DeleteCourseGrade(ctx, id)
}

func (r *mutationResolver) AddTuitionPayment(ctx context.Context, input model.AddTuitionPaymentInput) (*ent.TuitionPayment, error) {
	if !auth.IsAdmin(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.AddTuitionPayment(ctx, input)
}

func (r *mutationResolver) UpdateTuitionPayment(ctx context.Context, id uuid.UUID, input model.UpdateTuitionPaymentInput) (*ent.TuitionPayment, error) {
	if !auth.IsAdmin(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.UpdateTuitionPayment(ctx, id, input)
}

func (r *mutationResolver) DeleteTuitionPayment(ctx context.Context, id uuid.UUID) (bool, error) {
	if !auth.IsAdmin(ctx) {
		return false, auth.UnauthorizedErr
	}

	return true, r.s.DeleteTuitionPayment(ctx, id)
}

func (r *mutationResolver) AddAttendance(ctx context.Context, input model.AddAttendanceInput) (*ent.Attendance, error) {
	if !auth.IsAuthorized(ctx, user.RoleSuperAdmin, user.RoleSchoolAdmin, user.RoleTeacher) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.AddAttendance(ctx, input)
}

func (r *mutationResolver) UpdateAttendance(ctx context.Context, id uuid.UUID, input model.UpdateAttendanceInput) (*ent.Attendance, error) {
	if !auth.IsAuthorized(ctx, user.RoleSuperAdmin, user.RoleSchoolAdmin, user.RoleTeacher) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.UpdateAttendance(ctx, id, input)
}

func (r *mutationResolver) DeleteAttendance(ctx context.Context, id uuid.UUID) (bool, error) {
	if !auth.IsAuthorized(ctx, user.RoleSuperAdmin, user.RoleSchoolAdmin, user.RoleTeacher) {
		return false, auth.UnauthorizedErr
	}

	return true, r.s.DeleteAttendance(ctx, id)
}

func (r *mutationResolver) AddNotification(ctx context.Context, input model.AddNotificationInput) (*ent.Notification, error) {
	if !auth.IsAuthorized(ctx, user.RoleSuperAdmin, user.RoleSchoolAdmin, user.RoleTeacher) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.AddNotification(ctx, input)
}

func (r *mutationResolver) DeleteNotification(ctx context.Context, id uuid.UUID) (bool, error) {
	if !auth.IsAuthorized(ctx, user.RoleSuperAdmin, user.RoleSchoolAdmin, user.RoleTeacher) {
		return false, auth.UnauthorizedErr
	}

	return true, r.s.DeleteNotification(ctx, id)
}

func (r *queryResolver) Me(ctx context.Context) (*ent.User, error) {
	u, ok := auth.UserForContext(ctx)
	if !ok {
		return nil, auth.UnauthorizedErr
	}

	me, err := r.s.EC.User.Get(ctx, u.ID)
	if err != nil && ent.IsNotFound(err) {
		return nil, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: err.Error(), Extensions: map[string]interface{}{
				"code": "NOT_FOUND",
			}}
	}

	return me, err
}

func (r *queryResolver) School(ctx context.Context, id uuid.UUID) (*ent.School, error) {
	return r.s.EC.School.Get(ctx, id)
}

func (r *queryResolver) Schools(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.SchoolOrder, where *ent.SchoolWhereInput) (*ent.SchoolConnection, error) {
	if !auth.IsSuperAdmin(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.Schools(ctx, service.SchoolsOptions{
		After:   after,
		First:   first,
		Before:  before,
		Last:    last,
		OrderBy: orderBy,
		Where:   where,
	})
}

func (r *queryResolver) User(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	return r.s.EC.User.Get(ctx, id)
}

func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	if !auth.IsSuperAdmin(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.Users(ctx, service.UsersOptions{
		After:   after,
		First:   first,
		Before:  before,
		Last:    last,
		OrderBy: orderBy,
		Where:   where,
	})
}

func (r *queryResolver) Stage(ctx context.Context, id uuid.UUID) (*ent.Stage, error) {
	return r.s.EC.Stage.Get(ctx, id)
}

func (r *queryResolver) Stages(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.StageOrder, where *ent.StageWhereInput) (*ent.StageConnection, error) {
	return r.s.Stages(ctx, service.StagesOptions{
		After:   after,
		First:   first,
		Before:  before,
		Last:    last,
		OrderBy: orderBy,
		Where:   where,
	})
}

func (r *queryResolver) Messages(ctx context.Context, groupID uuid.UUID, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.MessageOrder, where *ent.MessageWhereInput) (*ent.MessageConnection, error) {
	u, ok := auth.UserForContext(ctx)
	if !ok {
		return nil, auth.UnauthorizedErr
	}

	if err := r.s.CheckAllowedToParticipateInChat(ctx, groupID, u.ID); err != nil {
		return nil, err
	}

	return r.s.Messages(ctx, groupID, service.MessagesOptions{After: after, First: first, Before: before, Last: last, OrderBy: orderBy, Where: where})
}

func (r *queryResolver) Group(ctx context.Context, id uuid.UUID) (*ent.Group, error) {
	return r.s.EC.Group.Get(ctx, id)
}

func (r *queryResolver) Groups(ctx context.Context, userID *uuid.UUID, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.GroupOrder, where *ent.GroupWhereInput) (*ent.GroupConnection, error) {
	u, ok := auth.UserForContext(ctx)
	if !ok {
		return nil, auth.UnauthorizedErr
	}

	if u.Role == user.RoleStudent || u.Role == user.RoleTeacher || userID == nil {
		userID = &u.ID
	}

	return r.s.Groups(ctx, service.GroupsOptions{
		UserID:  userID,
		After:   after,
		First:   first,
		Before:  before,
		Last:    last,
		OrderBy: orderBy,
		Where:   where,
	})
}

func (r *queryResolver) Class(ctx context.Context, id uuid.UUID) (*ent.Class, error) {
	return r.s.EC.Class.Get(ctx, id)
}

func (r *queryResolver) Classes(ctx context.Context, userID *uuid.UUID, stageID *uuid.UUID, schoolID *uuid.UUID, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ClassOrder, where *ent.ClassWhereInput) (*ent.ClassConnection, error) {
	u, ok := auth.UserForContext(ctx)
	if !ok {
		return nil, auth.UnauthorizedErr
	}

	// This step is just a precautionary measure and is probably unnecessary.
	// Make sure that an admin/teacher is bounded by their school and cannot see classes from other schools.
	if u.Role == user.RoleSchoolAdmin || u.Role == user.RoleTeacher {
		schID, err := r.s.EC.User.Query().Where(user.ID(u.ID)).QuerySchool().OnlyID(ctx)
		if err != nil {
			return nil, err
		}
		schoolID = &schID
	}

	if userID == nil {
		userID = &u.ID
	}

	return r.s.Classes(ctx, service.ClassesOptions{
		UserID:   userID,
		StageID:  stageID,
		SchoolID: schoolID,
		After:    after,
		First:    first,
		Before:   before,
		Last:     last,
		OrderBy:  orderBy,
		Where:    where,
	})
}

func (r *queryResolver) Assignment(ctx context.Context, id uuid.UUID) (*ent.Assignment, error) {
	return r.s.EC.Assignment.Get(ctx, id)
}

func (r *queryResolver) Assignments(ctx context.Context, userID *uuid.UUID, stageID *uuid.UUID, schoolID *uuid.UUID, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.AssignmentOrder, where *ent.AssignmentWhereInput) (*ent.AssignmentConnection, error) {
	u, ok := auth.UserForContext(ctx)
	if !ok {
		return nil, auth.UnauthorizedErr
	}

	if userID == nil {
		userID = &u.ID
	}

	return r.s.Assignments(ctx, service.AssignmentsOptions{
		UserID:  *userID,
		After:   after,
		First:   first,
		Before:  before,
		Last:    last,
		OrderBy: orderBy,
		Where:   where,
	})
}

func (r *queryResolver) AssignmentSubmissions(ctx context.Context, assignmentID *uuid.UUID, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.AssignmentSubmissionOrder, where *ent.AssignmentSubmissionWhereInput) (*ent.AssignmentSubmissionConnection, error) {
	if !auth.IsAuthorized(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.AssignmentSubmissions(ctx, *assignmentID, service.AssignmentSubmissionsOptions{
		After:   after,
		First:   first,
		Before:  before,
		Last:    last,
		OrderBy: orderBy,
		Where:   where,
	})
}

func (r *queryResolver) Schedule(ctx context.Context, stageID *uuid.UUID, weekday *time.Weekday) ([]*ent.Schedule, error) {
	u, ok := auth.UserForContext(ctx)
	if !ok {
		return nil, auth.UnauthorizedErr
	}

	return r.s.Schedule(ctx, service.ScheduleOptions{StageID: stageID, Weekday: weekday, UserID: u.ID})
}

func (r *queryResolver) CourseGrades(ctx context.Context, studentID *uuid.UUID, stageID *uuid.UUID, classID *uuid.UUID, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.CourseGradeOrder, where *ent.CourseGradeWhereInput) (*ent.CourseGradeConnection, error) {
	return r.s.CourseGrades(ctx, service.CourseGradesOptions{ClassID: classID, StudentID: studentID, StageID: stageID, After: after, First: first, Before: before, Last: last, OrderBy: orderBy, Where: where})
}

func (r *queryResolver) TuitionPayments(ctx context.Context, studentID *uuid.UUID, stageID *uuid.UUID, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TuitionPaymentOrder, where *ent.TuitionPaymentWhereInput) (*ent.TuitionPaymentConnection, error) {
	u, ok := auth.UserForContext(ctx)
	if !ok {
		return nil, auth.UnauthorizedErr
	}

	opts := service.TuitionPaymentsOptions{
		After:     after,
		First:     first,
		Before:    before,
		Last:      last,
		OrderBy:   orderBy,
		Where:     where,
		StageID:   stageID,
		StudentID: studentID,
	}

	switch u.Role {
	case user.RoleStudent:
		opts.StudentID = &u.ID

	case user.RoleSchoolAdmin, user.RoleSuperAdmin:
		if studentID == nil && stageID == nil {
			return nil, errors.New("stage id or student id is required")
		}

	default:
		return nil, auth.UnauthorizedErr
	}

	return r.s.TuitionPayments(ctx, opts)
}

func (r *queryResolver) Attendances(ctx context.Context, studentID *uuid.UUID, classID *uuid.UUID, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.AttendanceOrder, where *ent.AttendanceWhereInput) (*ent.AttendanceConnection, error) {
	u, ok := auth.UserForContext(ctx)
	if !ok {
		return nil, auth.UnauthorizedErr
	}

	if u.Role == user.RoleStudent {
		studentID = &u.ID
	}

	return r.s.Attendances(ctx, service.AttendancesOptions{
		After:     after,
		First:     first,
		Before:    before,
		Last:      last,
		StudentID: studentID,
		ClassID:   classID,
		OrderBy:   orderBy,
		Where:     where,
	})
}

func (r *queryResolver) Notifications(ctx context.Context, stageID *uuid.UUID, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.NotificationOrder, where *ent.NotificationWhereInput) (*ent.NotificationConnection, error) {
	u, ok := auth.UserForContext(ctx)
	if !ok {
		return nil, auth.UnauthorizedErr
	}

	usr, err := r.s.EC.User.Query().Where(user.ID(u.ID)).WithStage().WithSchool().Only(ctx)
	if err != nil {
		return nil, err
	}

	opts := service.NotificationsOptions{
		After:   after,
		First:   first,
		Before:  before,
		Last:    last,
		OrderBy: orderBy,
		Where:   where,
		StageID: stageID,
	}

	if usr.Role == user.RoleStudent {
		opts.StageID = &usr.Edges.Stage.ID
	}

	if usr.Role == user.RoleTeacher || usr.Role == user.RoleSchoolAdmin {
		opts.SchoolID = &usr.Edges.School.ID
	}

	return r.s.Notifications(ctx, opts)
}

func (r *schoolResolver) Users(ctx context.Context, obj *ent.School, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	return obj.QueryUsers().Where(user.DeletedAtIsNil()).Paginate(ctx, after, first, before, last, ent.WithUserOrder(orderBy), ent.WithUserFilter(where.Filter))
}

func (r *schoolResolver) Stages(ctx context.Context, obj *ent.School, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.StageOrder, where *ent.StageWhereInput) (*ent.StageConnection, error) {
	return obj.QueryStages().Where(stage.DeletedAtIsNil()).Paginate(ctx, after, first, before, last, ent.WithStageOrder(orderBy), ent.WithStageFilter(where.Filter))
}

func (r *stageResolver) Classes(ctx context.Context, obj *ent.Stage, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ClassOrder, where *ent.ClassWhereInput) (*ent.ClassConnection, error) {
	return obj.QueryClasses().Where(class.DeletedAtIsNil()).Paginate(ctx, after, first, before, last,
		ent.WithClassOrder(orderBy), ent.WithClassFilter(where.Filter))
}

func (r *stageResolver) Payments(ctx context.Context, obj *ent.Stage, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TuitionPaymentOrder, where *ent.TuitionPaymentWhereInput) (*ent.TuitionPaymentConnection, error) {
	return obj.QueryPayments().Paginate(ctx, after, first, before, last,
		ent.WithTuitionPaymentOrder(orderBy), ent.WithTuitionPaymentFilter(where.Filter))
}

func (r *stageResolver) Students(ctx context.Context, obj *ent.Stage, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	return obj.QueryStudents().Where(user.DeletedAtIsNil()).Paginate(ctx, after, first, before, last,
		ent.WithUserOrder(orderBy), ent.WithUserFilter(where.Filter))
}

func (r *subscriptionResolver) MessagePosted(ctx context.Context, groupID uuid.UUID) (<-chan *ent.Message, error) {
	u, ok := auth.UserForContext(ctx)
	if !ok {
		return nil, auth.UnauthorizedErr
	}

	return r.s.RegisterGroupObserver(ctx, groupID, u.ID)
}

func (r *userResolver) Messages(ctx context.Context, obj *ent.User, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.MessageOrder, where *ent.MessageWhereInput) (*ent.MessageConnection, error) {
	return obj.QueryMessages().Where(message.DeletedAtIsNil()).Paginate(ctx, after, first, before, last, ent.WithMessageOrder(orderBy), ent.WithMessageFilter(where.Filter))
}

func (r *userResolver) Groups(ctx context.Context, obj *ent.User, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.GroupOrder, where *ent.GroupWhereInput) (*ent.GroupConnection, error) {
	return obj.QueryGroups().Where(group.DeletedAtIsNil()).Paginate(ctx, after, first, before, last, ent.WithGroupOrder(orderBy), ent.WithGroupFilter(where.Filter))
}

func (r *userResolver) Classes(ctx context.Context, obj *ent.User, stageID *uuid.UUID, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ClassOrder, where *ent.ClassWhereInput) (*ent.ClassConnection, error) {
	return obj.QueryClasses().Where(class.DeletedAtIsNil()).Paginate(ctx, after, first, before, last, ent.WithClassOrder(orderBy), ent.WithClassFilter(where.Filter))
}

func (r *userResolver) AssignmentSubmissions(ctx context.Context, obj *ent.User, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.AssignmentSubmissionOrder, where *ent.AssignmentSubmissionWhereInput) (*ent.AssignmentSubmissionConnection, error) {
	return obj.QuerySubmissions().Paginate(ctx, after, first, before, last, ent.WithAssignmentSubmissionOrder(orderBy), ent.WithAssignmentSubmissionFilter(where.Filter))
}

func (r *userResolver) Payments(ctx context.Context, obj *ent.User, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TuitionPaymentOrder, where *ent.TuitionPaymentWhereInput) (*ent.TuitionPaymentConnection, error) {
	return obj.QueryPayments().Paginate(ctx, after, first, before, last,
		ent.WithTuitionPaymentOrder(orderBy), ent.WithTuitionPaymentFilter(where.Filter))
}

// Assignment returns generated.AssignmentResolver implementation.
func (r *Resolver) Assignment() generated.AssignmentResolver { return &assignmentResolver{r} }

// Class returns generated.ClassResolver implementation.
func (r *Resolver) Class() generated.ClassResolver { return &classResolver{r} }

// Group returns generated.GroupResolver implementation.
func (r *Resolver) Group() generated.GroupResolver { return &groupResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// School returns generated.SchoolResolver implementation.
func (r *Resolver) School() generated.SchoolResolver { return &schoolResolver{r} }

// Stage returns generated.StageResolver implementation.
func (r *Resolver) Stage() generated.StageResolver { return &stageResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type assignmentResolver struct{ *Resolver }
type classResolver struct{ *Resolver }
type groupResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type schoolResolver struct{ *Resolver }
type stageResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
