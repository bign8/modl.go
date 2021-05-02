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

	// EnterModl_pair is called when entering the modl_pair production.
	EnterModl_pair(c *Modl_pairContext)

	// EnterModl_value is called when entering the modl_value production.
	EnterModl_value(c *Modl_valueContext)

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

	// ExitModl_pair is called when exiting the modl_pair production.
	ExitModl_pair(c *Modl_pairContext)

	// ExitModl_value is called when exiting the modl_value production.
	ExitModl_value(c *Modl_valueContext)

	// ExitModl_primitive is called when exiting the modl_primitive production.
	ExitModl_primitive(c *Modl_primitiveContext)
}
