package dot

import (
	"fmt"

	"github.com/graphism/dot/ast"
	"github.com/mewkiz/pkg/errutil"
)

// check validates the semantics of the given DOT file.
func check(file *ast.File) error {
	for _, graph := range file.Graphs {
		// TODO: Check graph.ID for duplicates?
		if err := checkGraph(graph); err != nil {
			return errutil.Err(err)
		}
	}
	return nil
}

// check validates the semantics of the given graph.
func checkGraph(graph *ast.Graph) error {
	for _, stmt := range graph.Stmts {
		if err := checkStmt(graph, stmt); err != nil {
			return errutil.Err(err)
		}
	}
	return nil
}

// check validates the semantics of the given statement.
func checkStmt(graph *ast.Graph, stmt ast.Stmt) error {
	switch stmt := stmt.(type) {
	case *ast.NodeStmt:
		return checkNodeStmt(graph, stmt)
	case *ast.EdgeStmt:
		return checkEdgeStmt(graph, stmt)
	case *ast.AttrStmt:
		return checkAttrStmt(graph, stmt)
	case *ast.Attr:
		// TODO: Verify that the attribute is indeed of graph component kind.
		return checkAttr(graph, ast.KindGraph, stmt)
	case *ast.Subgraph:
		return checkSubgraph(graph, stmt)
	default:
		panic(fmt.Sprintf("support for statement of type %T not yet implemented", stmt))
	}
}

// checkNodeStmt validates the semantics of the given node statement.
func checkNodeStmt(graph *ast.Graph, stmt *ast.NodeStmt) error {
	if err := checkNode(graph, stmt.Node); err != nil {
		return errutil.Err(err)
	}
	for _, attr := range stmt.Attrs {
		// TODO: Verify that the attribute is indeed of node component kind.
		if err := checkAttr(graph, ast.KindNode, attr); err != nil {
			return errutil.Err(err)
		}
	}
	return nil
}

// checkEdgeStmt validates the semantics of the given edge statement.
func checkEdgeStmt(graph *ast.Graph, stmt *ast.EdgeStmt) error {
	// TODO: if graph.Strict, check for multi-edges.
	if err := checkVertex(graph, stmt.From); err != nil {
		return errutil.Err(err)
	}
	for _, attr := range stmt.Attrs {
		// TODO: Verify that the attribute is indeed of edge component kind.
		if err := checkAttr(graph, ast.KindEdge, attr); err != nil {
			return errutil.Err(err)
		}
	}
	return checkEdge(graph, stmt.From, stmt.To)
}

// checkEdge validates the semantics of the given edge.
func checkEdge(graph *ast.Graph, from ast.Vertex, to *ast.Edge) error {
	if !graph.Directed && to.Directed {
		return errutil.Newf("undirected graph %q contains directed edge from %q to %q", graph.ID, from, to.Vertex)
	}
	if err := checkVertex(graph, to.Vertex); err != nil {
		return errutil.Err(err)
	}
	if to.To != nil {
		return checkEdge(graph, to.Vertex, to.To)
	}
	return nil
}

// checkAttrStmt validates the semantics of the given attribute statement.
func checkAttrStmt(graph *ast.Graph, stmt *ast.AttrStmt) error {
	for _, attr := range stmt.Attrs {
		if err := checkAttr(graph, stmt.Kind, attr); err != nil {
			return errutil.Err(err)
		}
	}
	return nil
}

// checkAttr validates the semantics of the given attribute for the given
// component kind.
func checkAttr(graph *ast.Graph, kind ast.Kind, attr *ast.Attr) error {
	switch kind {
	case ast.KindGraph:
		// TODO: Validate key-value pairs for graphs.
		return nil
	case ast.KindNode:
		// TODO: Validate key-value pairs for nodes.
		return nil
	case ast.KindEdge:
		// TODO: Validate key-value pairs for edges.
		return nil
	default:
		panic(fmt.Sprintf("support for component kind %v not yet supported", kind))
	}
}

// checkSubgraph validates the semantics of the given subgraph.
func checkSubgraph(graph *ast.Graph, subgraph *ast.Subgraph) error {
	// TODO: Check subgraph.ID for duplicates?
	for _, stmt := range subgraph.Stmts {
		// TODO: Refine handling of subgraph statements?
		//    checkSubgraphStmt(graph, subgraph, stmt)
		if err := checkStmt(graph, stmt); err != nil {
			return errutil.Err(err)
		}
	}
	return nil
}

// checkVertex validates the semantics of the given vertex.
func checkVertex(graph *ast.Graph, vertex ast.Vertex) error {
	switch vertex := vertex.(type) {
	case *ast.Node:
		return checkNode(graph, vertex)
	case *ast.Subgraph:
		return checkSubgraph(graph, vertex)
	default:
		panic(fmt.Sprintf("support for vertex of type %T not yet supported", vertex))
	}
}

// checNode validates the semantics of the given node.
func checkNode(graph *ast.Graph, node *ast.Node) error {
	// TODO: Check node.ID for duplicates?
	// TODO: Validate node.Port.
	return nil
}
