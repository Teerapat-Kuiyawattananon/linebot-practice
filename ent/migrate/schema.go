// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CarsColumns holds the columns for the "cars" table.
	CarsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "model", Type: field.TypeString},
		{Name: "registered_at", Type: field.TypeTime},
		{Name: "price", Type: field.TypeInt, Default: 0},
		{Name: "image_path", Type: field.TypeString, Default: "https://digitalfinger.id/wp-content/uploads/2019/12/no-image-available-icon-6.png"},
		{Name: "line_user_cars", Type: field.TypeInt, Nullable: true},
	}
	// CarsTable holds the schema information for the "cars" table.
	CarsTable = &schema.Table{
		Name:       "cars",
		Columns:    CarsColumns,
		PrimaryKey: []*schema.Column{CarsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cars_line_users_cars",
				Columns:    []*schema.Column{CarsColumns[5]},
				RefColumns: []*schema.Column{LineUsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// CreditLatersColumns holds the columns for the "credit_laters" table.
	CreditLatersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "transaction_ref", Type: field.TypeString},
		{Name: "date", Type: field.TypeString, Default: "10/07/2023 23:08"},
		{Name: "branch", Type: field.TypeString, Default: "Center"},
		{Name: "amount", Type: field.TypeInt, Default: 0},
		{Name: "installment", Type: field.TypeInt, Default: 0},
		{Name: "detail", Type: field.TypeString, Default: "-"},
		{Name: "line_user_creditlaters", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// CreditLatersTable holds the schema information for the "credit_laters" table.
	CreditLatersTable = &schema.Table{
		Name:       "credit_laters",
		Columns:    CreditLatersColumns,
		PrimaryKey: []*schema.Column{CreditLatersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "credit_laters_line_users_creditlaters",
				Columns:    []*schema.Column{CreditLatersColumns[7]},
				RefColumns: []*schema.Column{LineUsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
	}
	// LineLogsColumns holds the columns for the "line_logs" table.
	LineLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "action", Type: field.TypeString, Default: "unknown"},
		{Name: "message", Type: field.TypeString, Default: "unknown"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "line_user_linelogs", Type: field.TypeInt, Nullable: true},
	}
	// LineLogsTable holds the schema information for the "line_logs" table.
	LineLogsTable = &schema.Table{
		Name:       "line_logs",
		Columns:    LineLogsColumns,
		PrimaryKey: []*schema.Column{LineLogsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "line_logs_line_users_linelogs",
				Columns:    []*schema.Column{LineLogsColumns[4]},
				RefColumns: []*schema.Column{LineUsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// LineUsersColumns holds the columns for the "line_users" table.
	LineUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeString, Unique: true},
		{Name: "displya_name", Type: field.TypeString},
		{Name: "registered_at", Type: field.TypeTime},
		{Name: "active", Type: field.TypeBool, Default: false},
	}
	// LineUsersTable holds the schema information for the "line_users" table.
	LineUsersTable = &schema.Table{
		Name:       "line_users",
		Columns:    LineUsersColumns,
		PrimaryKey: []*schema.Column{LineUsersColumns[0]},
	}
	// GroupLineusersColumns holds the columns for the "group_lineusers" table.
	GroupLineusersColumns = []*schema.Column{
		{Name: "group_id", Type: field.TypeInt},
		{Name: "line_user_id", Type: field.TypeInt},
	}
	// GroupLineusersTable holds the schema information for the "group_lineusers" table.
	GroupLineusersTable = &schema.Table{
		Name:       "group_lineusers",
		Columns:    GroupLineusersColumns,
		PrimaryKey: []*schema.Column{GroupLineusersColumns[0], GroupLineusersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_lineusers_group_id",
				Columns:    []*schema.Column{GroupLineusersColumns[0]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "group_lineusers_line_user_id",
				Columns:    []*schema.Column{GroupLineusersColumns[1]},
				RefColumns: []*schema.Column{LineUsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CarsTable,
		CreditLatersTable,
		GroupsTable,
		LineLogsTable,
		LineUsersTable,
		GroupLineusersTable,
	}
)

func init() {
	CarsTable.ForeignKeys[0].RefTable = LineUsersTable
	CreditLatersTable.ForeignKeys[0].RefTable = LineUsersTable
	LineLogsTable.ForeignKeys[0].RefTable = LineUsersTable
	GroupLineusersTable.ForeignKeys[0].RefTable = GroupsTable
	GroupLineusersTable.ForeignKeys[1].RefTable = LineUsersTable
}
