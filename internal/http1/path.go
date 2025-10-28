package http1

func PathNotFound(r *Request) *Response {
	return NewResponse404()
}
func PathMethodNotAllowed(r *Request) *Response {
	return NewResponse405()
}

func PathHello(r *Request) *Response {
	response := NewResponse(200, "<h1>Привет Хэндлер</h1>")
	response.AddHeaders(DefaultHeaders)
	return response
}

func PathAbout(r *Request) *Response {
	html := `
	<h1>About as</h1>
	<p>sdjkghsdj sgdl sldks sg</p>
	<p>sdjkghsdj sgdl slafaf fadks sg</p>
	<h2>About as</h2>
	<p>sdjkghsdj sgdl slfaskls;oaj;sadks sg</p>
	<p>sdjkghsasgko;gakogakpdj sgdl sldks sg</p>
	`
	Response := NewResponse(200, html)
	Response.AddHeaders(DefaultHeaders)
	return Response
}
