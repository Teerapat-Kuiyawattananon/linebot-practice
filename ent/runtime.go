// Code generated by ent, DO NOT EDIT.

package ent

import (
	"entdemo/ent/car"
	"entdemo/ent/creditlater"
	"entdemo/ent/group"
	"entdemo/ent/linelog"
	"entdemo/ent/lineuser"
	"entdemo/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	carFields := schema.Car{}.Fields()
	_ = carFields
	// carDescRegisteredAt is the schema descriptor for registered_at field.
	carDescRegisteredAt := carFields[1].Descriptor()
	// car.DefaultRegisteredAt holds the default value on creation for the registered_at field.
	car.DefaultRegisteredAt = carDescRegisteredAt.Default.(func() time.Time)
	// carDescPrice is the schema descriptor for price field.
	carDescPrice := carFields[2].Descriptor()
	// car.DefaultPrice holds the default value on creation for the price field.
	car.DefaultPrice = carDescPrice.Default.(int)
	// carDescImagePath is the schema descriptor for image_path field.
	carDescImagePath := carFields[3].Descriptor()
	// car.DefaultImagePath holds the default value on creation for the image_path field.
	car.DefaultImagePath = carDescImagePath.Default.(string)
	creditlaterFields := schema.CreditLater{}.Fields()
	_ = creditlaterFields
	// creditlaterDescTransactionRef is the schema descriptor for transaction_ref field.
	creditlaterDescTransactionRef := creditlaterFields[0].Descriptor()
	// creditlater.TransactionRefValidator is a validator for the "transaction_ref" field. It is called by the builders before save.
	creditlater.TransactionRefValidator = creditlaterDescTransactionRef.Validators[0].(func(string) error)
	// creditlaterDescDate is the schema descriptor for date field.
	creditlaterDescDate := creditlaterFields[1].Descriptor()
	// creditlater.DefaultDate holds the default value on creation for the date field.
	creditlater.DefaultDate = creditlaterDescDate.Default.(string)
	// creditlaterDescBranch is the schema descriptor for branch field.
	creditlaterDescBranch := creditlaterFields[2].Descriptor()
	// creditlater.DefaultBranch holds the default value on creation for the branch field.
	creditlater.DefaultBranch = creditlaterDescBranch.Default.(string)
	// creditlaterDescAmount is the schema descriptor for amount field.
	creditlaterDescAmount := creditlaterFields[3].Descriptor()
	// creditlater.DefaultAmount holds the default value on creation for the amount field.
	creditlater.DefaultAmount = creditlaterDescAmount.Default.(int)
	// creditlaterDescInstallment is the schema descriptor for installment field.
	creditlaterDescInstallment := creditlaterFields[4].Descriptor()
	// creditlater.DefaultInstallment holds the default value on creation for the installment field.
	creditlater.DefaultInstallment = creditlaterDescInstallment.Default.(int)
	// creditlaterDescDetail is the schema descriptor for detail field.
	creditlaterDescDetail := creditlaterFields[5].Descriptor()
	// creditlater.DefaultDetail holds the default value on creation for the detail field.
	creditlater.DefaultDetail = creditlaterDescDetail.Default.(string)
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescName is the schema descriptor for name field.
	groupDescName := groupFields[0].Descriptor()
	// group.NameValidator is a validator for the "name" field. It is called by the builders before save.
	group.NameValidator = groupDescName.Validators[0].(func(string) error)
	linelogFields := schema.LineLog{}.Fields()
	_ = linelogFields
	// linelogDescAction is the schema descriptor for action field.
	linelogDescAction := linelogFields[0].Descriptor()
	// linelog.DefaultAction holds the default value on creation for the action field.
	linelog.DefaultAction = linelogDescAction.Default.(string)
	// linelogDescMessage is the schema descriptor for message field.
	linelogDescMessage := linelogFields[1].Descriptor()
	// linelog.DefaultMessage holds the default value on creation for the message field.
	linelog.DefaultMessage = linelogDescMessage.Default.(string)
	// linelogDescCreatedAt is the schema descriptor for created_at field.
	linelogDescCreatedAt := linelogFields[2].Descriptor()
	// linelog.DefaultCreatedAt holds the default value on creation for the created_at field.
	linelog.DefaultCreatedAt = linelogDescCreatedAt.Default.(time.Time)
	lineuserFields := schema.LineUser{}.Fields()
	_ = lineuserFields
	// lineuserDescUserId is the schema descriptor for userId field.
	lineuserDescUserId := lineuserFields[0].Descriptor()
	// lineuser.UserIdValidator is a validator for the "userId" field. It is called by the builders before save.
	lineuser.UserIdValidator = lineuserDescUserId.Validators[0].(func(string) error)
	// lineuserDescRegisteredAt is the schema descriptor for registered_at field.
	lineuserDescRegisteredAt := lineuserFields[2].Descriptor()
	// lineuser.DefaultRegisteredAt holds the default value on creation for the registered_at field.
	lineuser.DefaultRegisteredAt = lineuserDescRegisteredAt.Default.(time.Time)
	// lineuserDescActive is the schema descriptor for active field.
	lineuserDescActive := lineuserFields[3].Descriptor()
	// lineuser.DefaultActive holds the default value on creation for the active field.
	lineuser.DefaultActive = lineuserDescActive.Default.(bool)
}
