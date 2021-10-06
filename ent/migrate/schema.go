// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AssignmentsColumns holds the columns for the "assignments" table.
	AssignmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "file", Type: field.TypeString, Nullable: true},
		{Name: "is_exam", Type: field.TypeBool, Default: false},
		{Name: "due_date", Type: field.TypeTime},
		{Name: "duration", Type: field.TypeInt64, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "class_assignments", Type: field.TypeUUID, Nullable: true},
	}
	// AssignmentsTable holds the schema information for the "assignments" table.
	AssignmentsTable = &schema.Table{
		Name:       "assignments",
		Columns:    AssignmentsColumns,
		PrimaryKey: []*schema.Column{AssignmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "assignments_classes_assignments",
				Columns:    []*schema.Column{AssignmentsColumns[10]},
				RefColumns: []*schema.Column{ClassesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// AssignmentSubmissionsColumns holds the columns for the "assignment_submissions" table.
	AssignmentSubmissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "files", Type: field.TypeJSON},
		{Name: "submitted_at", Type: field.TypeTime, Nullable: true},
		{Name: "assignment_submissions", Type: field.TypeUUID, Nullable: true},
		{Name: "user_submissions", Type: field.TypeUUID, Nullable: true},
	}
	// AssignmentSubmissionsTable holds the schema information for the "assignment_submissions" table.
	AssignmentSubmissionsTable = &schema.Table{
		Name:       "assignment_submissions",
		Columns:    AssignmentSubmissionsColumns,
		PrimaryKey: []*schema.Column{AssignmentSubmissionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "assignment_submissions_assignments_submissions",
				Columns:    []*schema.Column{AssignmentSubmissionsColumns[5]},
				RefColumns: []*schema.Column{AssignmentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "assignment_submissions_users_submissions",
				Columns:    []*schema.Column{AssignmentSubmissionsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "assignmentsubmission_user_submissions_assignment_submissions",
				Unique:  true,
				Columns: []*schema.Column{AssignmentSubmissionsColumns[6], AssignmentSubmissionsColumns[5]},
			},
			{
				Name:    "assignmentsubmission_assignment_submissions",
				Unique:  false,
				Columns: []*schema.Column{AssignmentSubmissionsColumns[5]},
			},
		},
	}
	// AttendancesColumns holds the columns for the "attendances" table.
	AttendancesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "date", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "date"}},
		{Name: "state", Type: field.TypeEnum, Enums: []string{"PRESENT", "ABSENT", "EXCUSED_ABSENCE", "SICK"}, Default: "PRESENT"},
		{Name: "class_attendances", Type: field.TypeUUID, Nullable: true},
		{Name: "user_attendances", Type: field.TypeUUID, Nullable: true},
	}
	// AttendancesTable holds the schema information for the "attendances" table.
	AttendancesTable = &schema.Table{
		Name:       "attendances",
		Columns:    AttendancesColumns,
		PrimaryKey: []*schema.Column{AttendancesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "attendances_classes_attendances",
				Columns:    []*schema.Column{AttendancesColumns[5]},
				RefColumns: []*schema.Column{ClassesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "attendances_users_attendances",
				Columns:    []*schema.Column{AttendancesColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "attendance_date_class_attendances_user_attendances",
				Unique:  false,
				Columns: []*schema.Column{AttendancesColumns[3], AttendancesColumns[5], AttendancesColumns[6]},
			},
			{
				Name:    "attendance_state",
				Unique:  false,
				Columns: []*schema.Column{AttendancesColumns[4]},
			},
			{
				Name:    "attendance_date",
				Unique:  false,
				Columns: []*schema.Column{AttendancesColumns[3]},
			},
		},
	}
	// ClassesColumns holds the columns for the "classes" table.
	ClassesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "active", Type: field.TypeBool, Default: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "stage_classes", Type: field.TypeUUID, Nullable: true},
		{Name: "user_classes", Type: field.TypeUUID, Nullable: true},
	}
	// ClassesTable holds the schema information for the "classes" table.
	ClassesTable = &schema.Table{
		Name:       "classes",
		Columns:    ClassesColumns,
		PrimaryKey: []*schema.Column{ClassesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "classes_stages_classes",
				Columns:    []*schema.Column{ClassesColumns[6]},
				RefColumns: []*schema.Column{StagesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "classes_users_classes",
				Columns:    []*schema.Column{ClassesColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "class_stage_classes",
				Unique:  false,
				Columns: []*schema.Column{ClassesColumns[6]},
			},
			{
				Name:    "class_user_classes",
				Unique:  false,
				Columns: []*schema.Column{ClassesColumns[7]},
			},
			{
				Name:    "class_active",
				Unique:  false,
				Columns: []*schema.Column{ClassesColumns[4]},
			},
			{
				Name:    "class_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{ClassesColumns[5]},
			},
		},
	}
	// CourseGradesColumns holds the columns for the "course_grades" table.
	CourseGradesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "course", Type: field.TypeEnum, Enums: []string{"FIRST", "SECOND"}},
		{Name: "activity_first", Type: field.TypeInt, Nullable: true},
		{Name: "activity_second", Type: field.TypeInt, Nullable: true},
		{Name: "written_first", Type: field.TypeInt, Nullable: true},
		{Name: "written_second", Type: field.TypeInt, Nullable: true},
		{Name: "course_final", Type: field.TypeInt, Nullable: true},
		{Name: "year", Type: field.TypeString},
		{Name: "class_course_grades", Type: field.TypeUUID, Nullable: true},
		{Name: "user_course_grades", Type: field.TypeUUID, Nullable: true},
	}
	// CourseGradesTable holds the schema information for the "course_grades" table.
	CourseGradesTable = &schema.Table{
		Name:       "course_grades",
		Columns:    CourseGradesColumns,
		PrimaryKey: []*schema.Column{CourseGradesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "course_grades_classes_course_grades",
				Columns:    []*schema.Column{CourseGradesColumns[10]},
				RefColumns: []*schema.Column{ClassesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "course_grades_users_course_grades",
				Columns:    []*schema.Column{CourseGradesColumns[11]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "coursegrade_course_year_user_course_grades_class_course_grades",
				Unique:  true,
				Columns: []*schema.Column{CourseGradesColumns[3], CourseGradesColumns[9], CourseGradesColumns[11], CourseGradesColumns[10]},
			},
		},
	}
	// GradesColumns holds the columns for the "grades" table.
	GradesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "exam_grade", Type: field.TypeInt},
		{Name: "assignment_grades", Type: field.TypeUUID, Nullable: true},
		{Name: "user_grades", Type: field.TypeUUID, Nullable: true},
	}
	// GradesTable holds the schema information for the "grades" table.
	GradesTable = &schema.Table{
		Name:       "grades",
		Columns:    GradesColumns,
		PrimaryKey: []*schema.Column{GradesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "grades_assignments_grades",
				Columns:    []*schema.Column{GradesColumns[4]},
				RefColumns: []*schema.Column{AssignmentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "grades_users_grades",
				Columns:    []*schema.Column{GradesColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "grade_user_grades_assignment_grades",
				Unique:  true,
				Columns: []*schema.Column{GradesColumns[5], GradesColumns[4]},
			},
		},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "group_type", Type: field.TypeEnum, Enums: []string{"PRIVATE", "SHARED"}, Default: "SHARED"},
		{Name: "active", Type: field.TypeBool, Default: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "class_group", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "groups_classes_group",
				Columns:    []*schema.Column{GroupsColumns[7]},
				RefColumns: []*schema.Column{ClassesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "group_class_group",
				Unique:  false,
				Columns: []*schema.Column{GroupsColumns[7]},
			},
			{
				Name:    "group_active",
				Unique:  false,
				Columns: []*schema.Column{GroupsColumns[5]},
			},
			{
				Name:    "group_group_type",
				Unique:  false,
				Columns: []*schema.Column{GroupsColumns[4]},
			},
			{
				Name:    "group_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{GroupsColumns[6]},
			},
		},
	}
	// MessagesColumns holds the columns for the "messages" table.
	MessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "content", Type: field.TypeString, Nullable: true},
		{Name: "attachment", Type: field.TypeString, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "group_messages", Type: field.TypeUUID, Nullable: true},
		{Name: "user_messages", Type: field.TypeUUID, Nullable: true},
	}
	// MessagesTable holds the schema information for the "messages" table.
	MessagesTable = &schema.Table{
		Name:       "messages",
		Columns:    MessagesColumns,
		PrimaryKey: []*schema.Column{MessagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "messages_groups_messages",
				Columns:    []*schema.Column{MessagesColumns[6]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "messages_users_messages",
				Columns:    []*schema.Column{MessagesColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "message_user_messages",
				Unique:  false,
				Columns: []*schema.Column{MessagesColumns[7]},
			},
			{
				Name:    "message_group_messages",
				Unique:  false,
				Columns: []*schema.Column{MessagesColumns[6]},
			},
			{
				Name:    "message_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{MessagesColumns[5]},
			},
		},
	}
	// NotificationsColumns holds the columns for the "notifications" table.
	NotificationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString},
		{Name: "body", Type: field.TypeString, Nullable: true},
		{Name: "image", Type: field.TypeString, Nullable: true},
		{Name: "route", Type: field.TypeString, Nullable: true, Size: 9},
		{Name: "color", Type: field.TypeString, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "stage_notifications", Type: field.TypeUUID, Nullable: true},
	}
	// NotificationsTable holds the schema information for the "notifications" table.
	NotificationsTable = &schema.Table{
		Name:       "notifications",
		Columns:    NotificationsColumns,
		PrimaryKey: []*schema.Column{NotificationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "notifications_stages_notifications",
				Columns:    []*schema.Column{NotificationsColumns[9]},
				RefColumns: []*schema.Column{StagesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "notification_stage_notifications",
				Unique:  false,
				Columns: []*schema.Column{NotificationsColumns[9]},
			},
			{
				Name:    "notification_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{NotificationsColumns[8]},
			},
		},
	}
	// SchedulesColumns holds the columns for the "schedules" table.
	SchedulesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "weekday", Type: field.TypeInt},
		{Name: "starts_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "time"}},
		{Name: "duration", Type: field.TypeInt64},
		{Name: "class_schedules", Type: field.TypeUUID, Nullable: true},
	}
	// SchedulesTable holds the schema information for the "schedules" table.
	SchedulesTable = &schema.Table{
		Name:       "schedules",
		Columns:    SchedulesColumns,
		PrimaryKey: []*schema.Column{SchedulesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "schedules_classes_schedules",
				Columns:    []*schema.Column{SchedulesColumns[4]},
				RefColumns: []*schema.Column{ClassesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "schedule_weekday_starts_at_class_schedules",
				Unique:  false,
				Columns: []*schema.Column{SchedulesColumns[1], SchedulesColumns[2], SchedulesColumns[4]},
			},
			{
				Name:    "schedule_weekday",
				Unique:  false,
				Columns: []*schema.Column{SchedulesColumns[1]},
			},
		},
	}
	// SchoolsColumns holds the columns for the "schools" table.
	SchoolsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "image", Type: field.TypeString},
		{Name: "directory", Type: field.TypeString},
		{Name: "active", Type: field.TypeBool, Default: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// SchoolsTable holds the schema information for the "schools" table.
	SchoolsTable = &schema.Table{
		Name:       "schools",
		Columns:    SchoolsColumns,
		PrimaryKey: []*schema.Column{SchoolsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "school_active",
				Unique:  false,
				Columns: []*schema.Column{SchoolsColumns[6]},
			},
			{
				Name:    "school_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{SchoolsColumns[7]},
			},
		},
	}
	// StagesColumns holds the columns for the "stages" table.
	StagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "tuition_amount", Type: field.TypeInt},
		{Name: "directory", Type: field.TypeString},
		{Name: "active", Type: field.TypeBool, Default: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "school_stages", Type: field.TypeUUID, Nullable: true},
	}
	// StagesTable holds the schema information for the "stages" table.
	StagesTable = &schema.Table{
		Name:       "stages",
		Columns:    StagesColumns,
		PrimaryKey: []*schema.Column{StagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "stages_schools_stages",
				Columns:    []*schema.Column{StagesColumns[8]},
				RefColumns: []*schema.Column{SchoolsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "stage_school_stages",
				Unique:  false,
				Columns: []*schema.Column{StagesColumns[8]},
			},
			{
				Name:    "stage_active",
				Unique:  false,
				Columns: []*schema.Column{StagesColumns[6]},
			},
			{
				Name:    "stage_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{StagesColumns[7]},
			},
		},
	}
	// TuitionPaymentsColumns holds the columns for the "tuition_payments" table.
	TuitionPaymentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "year", Type: field.TypeString},
		{Name: "paid_amount", Type: field.TypeInt},
		{Name: "stage_payments", Type: field.TypeUUID, Nullable: true},
		{Name: "user_payments", Type: field.TypeUUID, Nullable: true},
	}
	// TuitionPaymentsTable holds the schema information for the "tuition_payments" table.
	TuitionPaymentsTable = &schema.Table{
		Name:       "tuition_payments",
		Columns:    TuitionPaymentsColumns,
		PrimaryKey: []*schema.Column{TuitionPaymentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tuition_payments_stages_payments",
				Columns:    []*schema.Column{TuitionPaymentsColumns[5]},
				RefColumns: []*schema.Column{StagesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "tuition_payments_users_payments",
				Columns:    []*schema.Column{TuitionPaymentsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "tuitionpayment_user_payments",
				Unique:  false,
				Columns: []*schema.Column{TuitionPaymentsColumns[6]},
			},
			{
				Name:    "tuitionpayment_stage_payments",
				Unique:  false,
				Columns: []*schema.Column{TuitionPaymentsColumns[5]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "image", Type: field.TypeString, Nullable: true},
		{Name: "directory", Type: field.TypeString},
		{Name: "token_version", Type: field.TypeInt, Default: 0},
		{Name: "push_tokens", Type: field.TypeJSON, Nullable: true},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"SUPER_ADMIN", "SCHOOL_ADMIN", "TEACHER", "STUDENT"}, Default: "STUDENT"},
		{Name: "active", Type: field.TypeBool, Default: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "school_users", Type: field.TypeUUID, Nullable: true},
		{Name: "stage_students", Type: field.TypeUUID, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_schools_users",
				Columns:    []*schema.Column{UsersColumns[14]},
				RefColumns: []*schema.Column{SchoolsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "users_stages_students",
				Columns:    []*schema.Column{UsersColumns[15]},
				RefColumns: []*schema.Column{StagesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "user_stage_students",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[15]},
			},
			{
				Name:    "user_school_users",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[14]},
			},
			{
				Name:    "user_active",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[12]},
			},
			{
				Name:    "user_role",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[11]},
			},
			{
				Name:    "user_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[13]},
			},
		},
	}
	// UserGroupsColumns holds the columns for the "user_groups" table.
	UserGroupsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "group_id", Type: field.TypeUUID},
	}
	// UserGroupsTable holds the schema information for the "user_groups" table.
	UserGroupsTable = &schema.Table{
		Name:       "user_groups",
		Columns:    UserGroupsColumns,
		PrimaryKey: []*schema.Column{UserGroupsColumns[0], UserGroupsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_groups_user_id",
				Columns:    []*schema.Column{UserGroupsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_groups_group_id",
				Columns:    []*schema.Column{UserGroupsColumns[1]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AssignmentsTable,
		AssignmentSubmissionsTable,
		AttendancesTable,
		ClassesTable,
		CourseGradesTable,
		GradesTable,
		GroupsTable,
		MessagesTable,
		NotificationsTable,
		SchedulesTable,
		SchoolsTable,
		StagesTable,
		TuitionPaymentsTable,
		UsersTable,
		UserGroupsTable,
	}
)

func init() {
	AssignmentsTable.ForeignKeys[0].RefTable = ClassesTable
	AssignmentSubmissionsTable.ForeignKeys[0].RefTable = AssignmentsTable
	AssignmentSubmissionsTable.ForeignKeys[1].RefTable = UsersTable
	AttendancesTable.ForeignKeys[0].RefTable = ClassesTable
	AttendancesTable.ForeignKeys[1].RefTable = UsersTable
	ClassesTable.ForeignKeys[0].RefTable = StagesTable
	ClassesTable.ForeignKeys[1].RefTable = UsersTable
	CourseGradesTable.ForeignKeys[0].RefTable = ClassesTable
	CourseGradesTable.ForeignKeys[1].RefTable = UsersTable
	GradesTable.ForeignKeys[0].RefTable = AssignmentsTable
	GradesTable.ForeignKeys[1].RefTable = UsersTable
	GroupsTable.ForeignKeys[0].RefTable = ClassesTable
	MessagesTable.ForeignKeys[0].RefTable = GroupsTable
	MessagesTable.ForeignKeys[1].RefTable = UsersTable
	NotificationsTable.ForeignKeys[0].RefTable = StagesTable
	SchedulesTable.ForeignKeys[0].RefTable = ClassesTable
	StagesTable.ForeignKeys[0].RefTable = SchoolsTable
	TuitionPaymentsTable.ForeignKeys[0].RefTable = StagesTable
	TuitionPaymentsTable.ForeignKeys[1].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = SchoolsTable
	UsersTable.ForeignKeys[1].RefTable = StagesTable
	UserGroupsTable.ForeignKeys[0].RefTable = UsersTable
	UserGroupsTable.ForeignKeys[1].RefTable = GroupsTable
}
