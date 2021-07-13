package parser

type Tree struct {
	Attributes []*AttributeNode
	Types      []*TypeNode
	Server     *ServerNode
}

type AttributeNode struct {
	Name  string
	Value string
}

type TypeNode struct {
	Name       string
	Expression string
	Fields     []*TypeNode
}

type ServerNode struct {
	Attributes []*AttributeNode
	HTTP       *HTTPNode
	GRPC       *GRPCNode
}

type HTTPNode struct {
	Attributes  []*AttributeNode
	Middlewares []string
	Groups      []*HTTPGroupNode
	Routes      []*HTTPRouteNode
}

type HTTPGroupNode struct {
	Middlewares []string
	Relative    string
	Routes      []*HTTPRouteNode
}

type HTTPRouteNode struct {
	Methods     []string
	Middlewares []string
	Handler     *HandlerNode
}

type GRPCNode struct {
	Attributes  []*AttributeNode
	Middlewares []string
	Handlers    []*HandlerNode
}

type GRPCGroupNode struct {
	Middlewares []string
	Relative    string
	Handlers    []*HandlerNode
}

type GRPCRouteNode struct {
	Middlewares []string
	Handler     *HandlerNode
}

type HandlerNode struct {
	Name   string
	Input  string
	Output string
}
