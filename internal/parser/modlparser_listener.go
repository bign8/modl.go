// Code generated from grammar/antlr4/MODLParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // MODLParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MODLParserListener is a complete listener for a parse tree produced by MODLParser.
type MODLParserListener interface {
	antlr.ParseTreeListener

	// EnterModl is called when entering the modl production.
	EnterModl(c *ModlContext)

	// EnterModl_structure is called when entering the modl_structure production.
	EnterModl_structure(c *Modl_structureContext)

	// EnterModl_map is called when entering the modl_map production.
	EnterModl_map(c *Modl_mapContext)

	// EnterModl_array is called when entering the modl_array production.
	EnterModl_array(c *Modl_arrayContext)

	// EnterModl_nb_array is called when entering the modl_nb_array production.
	EnterModl_nb_array(c *Modl_nb_arrayContext)

	// EnterModl_pair is called when entering the modl_pair production.
	EnterModl_pair(c *Modl_pairContext)

	// EnterModl_value_item is called when entering the modl_value_item production.
	EnterModl_value_item(c *Modl_value_itemContext)

	// EnterModl_top_level_conditional is called when entering the modl_top_level_conditional production.
	EnterModl_top_level_conditional(c *Modl_top_level_conditionalContext)

	// EnterModl_top_level_conditional_return is called when entering the modl_top_level_conditional_return production.
	EnterModl_top_level_conditional_return(c *Modl_top_level_conditional_returnContext)

	// EnterModl_map_conditional is called when entering the modl_map_conditional production.
	EnterModl_map_conditional(c *Modl_map_conditionalContext)

	// EnterModl_map_conditional_return is called when entering the modl_map_conditional_return production.
	EnterModl_map_conditional_return(c *Modl_map_conditional_returnContext)

	// EnterModl_map_item is called when entering the modl_map_item production.
	EnterModl_map_item(c *Modl_map_itemContext)

	// EnterModl_array_conditional is called when entering the modl_array_conditional production.
	EnterModl_array_conditional(c *Modl_array_conditionalContext)

	// EnterModl_array_conditional_return is called when entering the modl_array_conditional_return production.
	EnterModl_array_conditional_return(c *Modl_array_conditional_returnContext)

	// EnterModl_array_item is called when entering the modl_array_item production.
	EnterModl_array_item(c *Modl_array_itemContext)

	// EnterModl_value_conditional is called when entering the modl_value_conditional production.
	EnterModl_value_conditional(c *Modl_value_conditionalContext)

	// EnterModl_value_conditional_return is called when entering the modl_value_conditional_return production.
	EnterModl_value_conditional_return(c *Modl_value_conditional_returnContext)

	// EnterModl_condition_test is called when entering the modl_condition_test production.
	EnterModl_condition_test(c *Modl_condition_testContext)

	// EnterModl_operator is called when entering the modl_operator production.
	EnterModl_operator(c *Modl_operatorContext)

	// EnterModl_condition is called when entering the modl_condition production.
	EnterModl_condition(c *Modl_conditionContext)

	// EnterModl_condition_group is called when entering the modl_condition_group production.
	EnterModl_condition_group(c *Modl_condition_groupContext)

	// EnterModl_value is called when entering the modl_value production.
	EnterModl_value(c *Modl_valueContext)

	// EnterModl_array_value_item is called when entering the modl_array_value_item production.
	EnterModl_array_value_item(c *Modl_array_value_itemContext)

	// EnterModl_primitive is called when entering the modl_primitive production.
	EnterModl_primitive(c *Modl_primitiveContext)

	// ExitModl is called when exiting the modl production.
	ExitModl(c *ModlContext)

	// ExitModl_structure is called when exiting the modl_structure production.
	ExitModl_structure(c *Modl_structureContext)

	// ExitModl_map is called when exiting the modl_map production.
	ExitModl_map(c *Modl_mapContext)

	// ExitModl_array is called when exiting the modl_array production.
	ExitModl_array(c *Modl_arrayContext)

	// ExitModl_nb_array is called when exiting the modl_nb_array production.
	ExitModl_nb_array(c *Modl_nb_arrayContext)

	// ExitModl_pair is called when exiting the modl_pair production.
	ExitModl_pair(c *Modl_pairContext)

	// ExitModl_value_item is called when exiting the modl_value_item production.
	ExitModl_value_item(c *Modl_value_itemContext)

	// ExitModl_top_level_conditional is called when exiting the modl_top_level_conditional production.
	ExitModl_top_level_conditional(c *Modl_top_level_conditionalContext)

	// ExitModl_top_level_conditional_return is called when exiting the modl_top_level_conditional_return production.
	ExitModl_top_level_conditional_return(c *Modl_top_level_conditional_returnContext)

	// ExitModl_map_conditional is called when exiting the modl_map_conditional production.
	ExitModl_map_conditional(c *Modl_map_conditionalContext)

	// ExitModl_map_conditional_return is called when exiting the modl_map_conditional_return production.
	ExitModl_map_conditional_return(c *Modl_map_conditional_returnContext)

	// ExitModl_map_item is called when exiting the modl_map_item production.
	ExitModl_map_item(c *Modl_map_itemContext)

	// ExitModl_array_conditional is called when exiting the modl_array_conditional production.
	ExitModl_array_conditional(c *Modl_array_conditionalContext)

	// ExitModl_array_conditional_return is called when exiting the modl_array_conditional_return production.
	ExitModl_array_conditional_return(c *Modl_array_conditional_returnContext)

	// ExitModl_array_item is called when exiting the modl_array_item production.
	ExitModl_array_item(c *Modl_array_itemContext)

	// ExitModl_value_conditional is called when exiting the modl_value_conditional production.
	ExitModl_value_conditional(c *Modl_value_conditionalContext)

	// ExitModl_value_conditional_return is called when exiting the modl_value_conditional_return production.
	ExitModl_value_conditional_return(c *Modl_value_conditional_returnContext)

	// ExitModl_condition_test is called when exiting the modl_condition_test production.
	ExitModl_condition_test(c *Modl_condition_testContext)

	// ExitModl_operator is called when exiting the modl_operator production.
	ExitModl_operator(c *Modl_operatorContext)

	// ExitModl_condition is called when exiting the modl_condition production.
	ExitModl_condition(c *Modl_conditionContext)

	// ExitModl_condition_group is called when exiting the modl_condition_group production.
	ExitModl_condition_group(c *Modl_condition_groupContext)

	// ExitModl_value is called when exiting the modl_value production.
	ExitModl_value(c *Modl_valueContext)

	// ExitModl_array_value_item is called when exiting the modl_array_value_item production.
	ExitModl_array_value_item(c *Modl_array_value_itemContext)

	// ExitModl_primitive is called when exiting the modl_primitive production.
	ExitModl_primitive(c *Modl_primitiveContext)
}
