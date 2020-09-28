// Code generated from grammar/antlr4/MODLParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // MODLParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMODLParserListener is a complete listener for a parse tree produced by MODLParser.
type BaseMODLParserListener struct{}

var _ MODLParserListener = &BaseMODLParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMODLParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMODLParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMODLParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMODLParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterModl is called when production modl is entered.
func (s *BaseMODLParserListener) EnterModl(ctx *ModlContext) {}

// ExitModl is called when production modl is exited.
func (s *BaseMODLParserListener) ExitModl(ctx *ModlContext) {}

// EnterModl_structure is called when production modl_structure is entered.
func (s *BaseMODLParserListener) EnterModl_structure(ctx *Modl_structureContext) {}

// ExitModl_structure is called when production modl_structure is exited.
func (s *BaseMODLParserListener) ExitModl_structure(ctx *Modl_structureContext) {}

// EnterModl_map is called when production modl_map is entered.
func (s *BaseMODLParserListener) EnterModl_map(ctx *Modl_mapContext) {}

// ExitModl_map is called when production modl_map is exited.
func (s *BaseMODLParserListener) ExitModl_map(ctx *Modl_mapContext) {}

// EnterModl_array is called when production modl_array is entered.
func (s *BaseMODLParserListener) EnterModl_array(ctx *Modl_arrayContext) {}

// ExitModl_array is called when production modl_array is exited.
func (s *BaseMODLParserListener) ExitModl_array(ctx *Modl_arrayContext) {}

// EnterModl_nb_array is called when production modl_nb_array is entered.
func (s *BaseMODLParserListener) EnterModl_nb_array(ctx *Modl_nb_arrayContext) {}

// ExitModl_nb_array is called when production modl_nb_array is exited.
func (s *BaseMODLParserListener) ExitModl_nb_array(ctx *Modl_nb_arrayContext) {}

// EnterModl_pair is called when production modl_pair is entered.
func (s *BaseMODLParserListener) EnterModl_pair(ctx *Modl_pairContext) {}

// ExitModl_pair is called when production modl_pair is exited.
func (s *BaseMODLParserListener) ExitModl_pair(ctx *Modl_pairContext) {}

// EnterModl_value_item is called when production modl_value_item is entered.
func (s *BaseMODLParserListener) EnterModl_value_item(ctx *Modl_value_itemContext) {}

// ExitModl_value_item is called when production modl_value_item is exited.
func (s *BaseMODLParserListener) ExitModl_value_item(ctx *Modl_value_itemContext) {}

// EnterModl_top_level_conditional is called when production modl_top_level_conditional is entered.
func (s *BaseMODLParserListener) EnterModl_top_level_conditional(ctx *Modl_top_level_conditionalContext) {
}

// ExitModl_top_level_conditional is called when production modl_top_level_conditional is exited.
func (s *BaseMODLParserListener) ExitModl_top_level_conditional(ctx *Modl_top_level_conditionalContext) {
}

// EnterModl_top_level_conditional_return is called when production modl_top_level_conditional_return is entered.
func (s *BaseMODLParserListener) EnterModl_top_level_conditional_return(ctx *Modl_top_level_conditional_returnContext) {
}

// ExitModl_top_level_conditional_return is called when production modl_top_level_conditional_return is exited.
func (s *BaseMODLParserListener) ExitModl_top_level_conditional_return(ctx *Modl_top_level_conditional_returnContext) {
}

// EnterModl_map_conditional is called when production modl_map_conditional is entered.
func (s *BaseMODLParserListener) EnterModl_map_conditional(ctx *Modl_map_conditionalContext) {}

// ExitModl_map_conditional is called when production modl_map_conditional is exited.
func (s *BaseMODLParserListener) ExitModl_map_conditional(ctx *Modl_map_conditionalContext) {}

// EnterModl_map_conditional_return is called when production modl_map_conditional_return is entered.
func (s *BaseMODLParserListener) EnterModl_map_conditional_return(ctx *Modl_map_conditional_returnContext) {
}

// ExitModl_map_conditional_return is called when production modl_map_conditional_return is exited.
func (s *BaseMODLParserListener) ExitModl_map_conditional_return(ctx *Modl_map_conditional_returnContext) {
}

// EnterModl_map_item is called when production modl_map_item is entered.
func (s *BaseMODLParserListener) EnterModl_map_item(ctx *Modl_map_itemContext) {}

// ExitModl_map_item is called when production modl_map_item is exited.
func (s *BaseMODLParserListener) ExitModl_map_item(ctx *Modl_map_itemContext) {}

// EnterModl_array_conditional is called when production modl_array_conditional is entered.
func (s *BaseMODLParserListener) EnterModl_array_conditional(ctx *Modl_array_conditionalContext) {}

// ExitModl_array_conditional is called when production modl_array_conditional is exited.
func (s *BaseMODLParserListener) ExitModl_array_conditional(ctx *Modl_array_conditionalContext) {}

// EnterModl_array_conditional_return is called when production modl_array_conditional_return is entered.
func (s *BaseMODLParserListener) EnterModl_array_conditional_return(ctx *Modl_array_conditional_returnContext) {
}

// ExitModl_array_conditional_return is called when production modl_array_conditional_return is exited.
func (s *BaseMODLParserListener) ExitModl_array_conditional_return(ctx *Modl_array_conditional_returnContext) {
}

// EnterModl_array_item is called when production modl_array_item is entered.
func (s *BaseMODLParserListener) EnterModl_array_item(ctx *Modl_array_itemContext) {}

// ExitModl_array_item is called when production modl_array_item is exited.
func (s *BaseMODLParserListener) ExitModl_array_item(ctx *Modl_array_itemContext) {}

// EnterModl_value_conditional is called when production modl_value_conditional is entered.
func (s *BaseMODLParserListener) EnterModl_value_conditional(ctx *Modl_value_conditionalContext) {}

// ExitModl_value_conditional is called when production modl_value_conditional is exited.
func (s *BaseMODLParserListener) ExitModl_value_conditional(ctx *Modl_value_conditionalContext) {}

// EnterModl_value_conditional_return is called when production modl_value_conditional_return is entered.
func (s *BaseMODLParserListener) EnterModl_value_conditional_return(ctx *Modl_value_conditional_returnContext) {
}

// ExitModl_value_conditional_return is called when production modl_value_conditional_return is exited.
func (s *BaseMODLParserListener) ExitModl_value_conditional_return(ctx *Modl_value_conditional_returnContext) {
}

// EnterModl_condition_test is called when production modl_condition_test is entered.
func (s *BaseMODLParserListener) EnterModl_condition_test(ctx *Modl_condition_testContext) {}

// ExitModl_condition_test is called when production modl_condition_test is exited.
func (s *BaseMODLParserListener) ExitModl_condition_test(ctx *Modl_condition_testContext) {}

// EnterModl_operator is called when production modl_operator is entered.
func (s *BaseMODLParserListener) EnterModl_operator(ctx *Modl_operatorContext) {}

// ExitModl_operator is called when production modl_operator is exited.
func (s *BaseMODLParserListener) ExitModl_operator(ctx *Modl_operatorContext) {}

// EnterModl_condition is called when production modl_condition is entered.
func (s *BaseMODLParserListener) EnterModl_condition(ctx *Modl_conditionContext) {}

// ExitModl_condition is called when production modl_condition is exited.
func (s *BaseMODLParserListener) ExitModl_condition(ctx *Modl_conditionContext) {}

// EnterModl_condition_group is called when production modl_condition_group is entered.
func (s *BaseMODLParserListener) EnterModl_condition_group(ctx *Modl_condition_groupContext) {}

// ExitModl_condition_group is called when production modl_condition_group is exited.
func (s *BaseMODLParserListener) ExitModl_condition_group(ctx *Modl_condition_groupContext) {}

// EnterModl_value is called when production modl_value is entered.
func (s *BaseMODLParserListener) EnterModl_value(ctx *Modl_valueContext) {}

// ExitModl_value is called when production modl_value is exited.
func (s *BaseMODLParserListener) ExitModl_value(ctx *Modl_valueContext) {}

// EnterModl_array_value_item is called when production modl_array_value_item is entered.
func (s *BaseMODLParserListener) EnterModl_array_value_item(ctx *Modl_array_value_itemContext) {}

// ExitModl_array_value_item is called when production modl_array_value_item is exited.
func (s *BaseMODLParserListener) ExitModl_array_value_item(ctx *Modl_array_value_itemContext) {}

// EnterModl_primitive is called when production modl_primitive is entered.
func (s *BaseMODLParserListener) EnterModl_primitive(ctx *Modl_primitiveContext) {}

// ExitModl_primitive is called when production modl_primitive is exited.
func (s *BaseMODLParserListener) ExitModl_primitive(ctx *Modl_primitiveContext) {}
