// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonD2b7633eDecodeAForumInternalModels(in *jlexer.Lexer, out *Vote) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int(in.Int())
		case "nickname":
			out.Nickname = string(in.String())
		case "voice":
			out.Voice = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeAForumInternalModels(out *jwriter.Writer, in Vote) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"nickname\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Nickname))
	}
	{
		const prefix string = ",\"voice\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Voice))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Vote) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeAForumInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Vote) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeAForumInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Vote) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeAForumInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Vote) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeAForumInternalModels(l, v)
}
func easyjsonD2b7633eDecodeAForumInternalModels1(in *jlexer.Lexer, out *Users) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Users, 0, 8)
			} else {
				*out = Users{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 *User
			if in.IsNull() {
				in.Skip()
				v1 = nil
			} else {
				if v1 == nil {
					v1 = new(User)
				}
				(*v1).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeAForumInternalModels1(out *jwriter.Writer, in Users) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			if v3 == nil {
				out.RawString("null")
			} else {
				(*v3).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Users) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeAForumInternalModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Users) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeAForumInternalModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Users) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeAForumInternalModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Users) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeAForumInternalModels1(l, v)
}
func easyjsonD2b7633eDecodeAForumInternalModels2(in *jlexer.Lexer, out *User) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "about":
			out.About = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "fullname":
			out.Fullname = string(in.String())
		case "nickname":
			out.Nickname = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeAForumInternalModels2(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"about\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.About))
	}
	{
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"fullname\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Fullname))
	}
	{
		const prefix string = ",\"nickname\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Nickname))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeAForumInternalModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeAForumInternalModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeAForumInternalModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeAForumInternalModels2(l, v)
}
func easyjsonD2b7633eDecodeAForumInternalModels3(in *jlexer.Lexer, out *Threads) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Threads, 0, 8)
			} else {
				*out = Threads{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v4 *Thread
			if in.IsNull() {
				in.Skip()
				v4 = nil
			} else {
				if v4 == nil {
					v4 = new(Thread)
				}
				(*v4).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v4)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeAForumInternalModels3(out *jwriter.Writer, in Threads) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v5, v6 := range in {
			if v5 > 0 {
				out.RawByte(',')
			}
			if v6 == nil {
				out.RawString("null")
			} else {
				(*v6).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Threads) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeAForumInternalModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Threads) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeAForumInternalModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Threads) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeAForumInternalModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Threads) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeAForumInternalModels3(l, v)
}
func easyjsonD2b7633eDecodeAForumInternalModels4(in *jlexer.Lexer, out *Thread) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int(in.Int())
		case "author":
			out.Author = string(in.String())
		case "created":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Created).UnmarshalJSON(data))
			}
		case "forum":
			out.Forum = string(in.String())
		case "message":
			out.Message = string(in.String())
		case "slug":
			out.Slug = string(in.String())
		case "title":
			out.Title = string(in.String())
		case "votes":
			out.Votes = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeAForumInternalModels4(out *jwriter.Writer, in Thread) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"author\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Author))
	}
	{
		const prefix string = ",\"created\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Created).MarshalJSON())
	}
	{
		const prefix string = ",\"forum\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Forum))
	}
	{
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"slug\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Slug))
	}
	{
		const prefix string = ",\"title\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"votes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Votes))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Thread) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeAForumInternalModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Thread) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeAForumInternalModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Thread) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeAForumInternalModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Thread) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeAForumInternalModels4(l, v)
}
func easyjsonD2b7633eDecodeAForumInternalModels5(in *jlexer.Lexer, out *Status) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "forum":
			out.Forum = int(in.Int())
		case "post":
			out.Post = int(in.Int())
		case "thread":
			out.Thread = int(in.Int())
		case "user":
			out.User = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeAForumInternalModels5(out *jwriter.Writer, in Status) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"forum\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Forum))
	}
	{
		const prefix string = ",\"post\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Post))
	}
	{
		const prefix string = ",\"thread\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Thread))
	}
	{
		const prefix string = ",\"user\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.User))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Status) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeAForumInternalModels5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Status) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeAForumInternalModels5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Status) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeAForumInternalModels5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Status) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeAForumInternalModels5(l, v)
}
func easyjsonD2b7633eDecodeAForumInternalModels6(in *jlexer.Lexer, out *Posts) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Posts, 0, 8)
			} else {
				*out = Posts{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v7 *Post
			if in.IsNull() {
				in.Skip()
				v7 = nil
			} else {
				if v7 == nil {
					v7 = new(Post)
				}
				(*v7).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v7)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeAForumInternalModels6(out *jwriter.Writer, in Posts) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v8, v9 := range in {
			if v8 > 0 {
				out.RawByte(',')
			}
			if v9 == nil {
				out.RawString("null")
			} else {
				(*v9).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Posts) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeAForumInternalModels6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Posts) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeAForumInternalModels6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Posts) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeAForumInternalModels6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Posts) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeAForumInternalModels6(l, v)
}
func easyjsonD2b7633eDecodeAForumInternalModels7(in *jlexer.Lexer, out *PostQuery) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Desc":
			if in.IsNull() {
				in.Skip()
				out.Desc = nil
			} else {
				if out.Desc == nil {
					out.Desc = new(bool)
				}
				*out.Desc = bool(in.Bool())
			}
		case "Limit":
			if in.IsNull() {
				in.Skip()
				out.Limit = nil
			} else {
				if out.Limit == nil {
					out.Limit = new(int)
				}
				*out.Limit = int(in.Int())
			}
		case "Since":
			if in.IsNull() {
				in.Skip()
				out.Since = nil
			} else {
				if out.Since == nil {
					out.Since = new(int64)
				}
				*out.Since = int64(in.Int64())
			}
		case "SlugOrID":
			out.SlugOrID = string(in.String())
		case "Sort":
			out.Sort = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeAForumInternalModels7(out *jwriter.Writer, in PostQuery) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Desc\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Desc == nil {
			out.RawString("null")
		} else {
			out.Bool(bool(*in.Desc))
		}
	}
	{
		const prefix string = ",\"Limit\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Limit == nil {
			out.RawString("null")
		} else {
			out.Int(int(*in.Limit))
		}
	}
	{
		const prefix string = ",\"Since\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Since == nil {
			out.RawString("null")
		} else {
			out.Int64(int64(*in.Since))
		}
	}
	{
		const prefix string = ",\"SlugOrID\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SlugOrID))
	}
	{
		const prefix string = ",\"Sort\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Sort))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostQuery) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeAForumInternalModels7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostQuery) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeAForumInternalModels7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostQuery) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeAForumInternalModels7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostQuery) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeAForumInternalModels7(l, v)
}
func easyjsonD2b7633eDecodeAForumInternalModels8(in *jlexer.Lexer, out *PostFull) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "author":
			if in.IsNull() {
				in.Skip()
				out.Author = nil
			} else {
				if out.Author == nil {
					out.Author = new(User)
				}
				(*out.Author).UnmarshalEasyJSON(in)
			}
		case "forum":
			if in.IsNull() {
				in.Skip()
				out.Forum = nil
			} else {
				if out.Forum == nil {
					out.Forum = new(Forum)
				}
				(*out.Forum).UnmarshalEasyJSON(in)
			}
		case "post":
			if in.IsNull() {
				in.Skip()
				out.Post = nil
			} else {
				if out.Post == nil {
					out.Post = new(Post)
				}
				(*out.Post).UnmarshalEasyJSON(in)
			}
		case "thread":
			if in.IsNull() {
				in.Skip()
				out.Thread = nil
			} else {
				if out.Thread == nil {
					out.Thread = new(Thread)
				}
				(*out.Thread).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeAForumInternalModels8(out *jwriter.Writer, in PostFull) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Author != nil {
		const prefix string = ",\"author\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Author).MarshalEasyJSON(out)
	}
	if in.Forum != nil {
		const prefix string = ",\"forum\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Forum).MarshalEasyJSON(out)
	}
	if in.Post != nil {
		const prefix string = ",\"post\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Post).MarshalEasyJSON(out)
	}
	if in.Thread != nil {
		const prefix string = ",\"thread\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Thread).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostFull) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeAForumInternalModels8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostFull) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeAForumInternalModels8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostFull) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeAForumInternalModels8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostFull) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeAForumInternalModels8(l, v)
}
func easyjsonD2b7633eDecodeAForumInternalModels9(in *jlexer.Lexer, out *Post) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int64(in.Int64())
		case "author":
			out.Author = string(in.String())
		case "created":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Created).UnmarshalJSON(data))
			}
		case "forum":
			out.Forum = string(in.String())
		case "isEdited":
			out.IsEdited = bool(in.Bool())
		case "message":
			out.Message = string(in.String())
		case "parent":
			out.Parent = int64(in.Int64())
		case "thread":
			out.Thread = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeAForumInternalModels9(out *jwriter.Writer, in Post) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != 0 {
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"author\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Author))
	}
	{
		const prefix string = ",\"created\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Created).MarshalJSON())
	}
	if in.Forum != "" {
		const prefix string = ",\"forum\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Forum))
	}
	if in.IsEdited {
		const prefix string = ",\"isEdited\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IsEdited))
	}
	{
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	if in.Parent != 0 {
		const prefix string = ",\"parent\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Parent))
	}
	if in.Thread != 0 {
		const prefix string = ",\"thread\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Thread))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Post) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeAForumInternalModels9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Post) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeAForumInternalModels9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Post) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeAForumInternalModels9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Post) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeAForumInternalModels9(l, v)
}
func easyjsonD2b7633eDecodeAForumInternalModels10(in *jlexer.Lexer, out *Forums) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Forums, 0, 8)
			} else {
				*out = Forums{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v10 *Forum
			if in.IsNull() {
				in.Skip()
				v10 = nil
			} else {
				if v10 == nil {
					v10 = new(Forum)
				}
				(*v10).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v10)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeAForumInternalModels10(out *jwriter.Writer, in Forums) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v11, v12 := range in {
			if v11 > 0 {
				out.RawByte(',')
			}
			if v12 == nil {
				out.RawString("null")
			} else {
				(*v12).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Forums) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeAForumInternalModels10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Forums) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeAForumInternalModels10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Forums) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeAForumInternalModels10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Forums) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeAForumInternalModels10(l, v)
}
func easyjsonD2b7633eDecodeAForumInternalModels11(in *jlexer.Lexer, out *ForumQuery) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Desc":
			if in.IsNull() {
				in.Skip()
				out.Desc = nil
			} else {
				if out.Desc == nil {
					out.Desc = new(bool)
				}
				*out.Desc = bool(in.Bool())
			}
		case "Limit":
			if in.IsNull() {
				in.Skip()
				out.Limit = nil
			} else {
				if out.Limit == nil {
					out.Limit = new(int)
				}
				*out.Limit = int(in.Int())
			}
		case "Since":
			if in.IsNull() {
				in.Skip()
				out.Since = nil
			} else {
				if out.Since == nil {
					out.Since = new(string)
				}
				*out.Since = string(in.String())
			}
		case "Slug":
			out.Slug = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeAForumInternalModels11(out *jwriter.Writer, in ForumQuery) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Desc\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Desc == nil {
			out.RawString("null")
		} else {
			out.Bool(bool(*in.Desc))
		}
	}
	{
		const prefix string = ",\"Limit\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Limit == nil {
			out.RawString("null")
		} else {
			out.Int(int(*in.Limit))
		}
	}
	{
		const prefix string = ",\"Since\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Since == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Since))
		}
	}
	{
		const prefix string = ",\"Slug\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Slug))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ForumQuery) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeAForumInternalModels11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ForumQuery) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeAForumInternalModels11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ForumQuery) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeAForumInternalModels11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ForumQuery) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeAForumInternalModels11(l, v)
}
func easyjsonD2b7633eDecodeAForumInternalModels12(in *jlexer.Lexer, out *Forum) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "posts":
			out.Posts = int(in.Int())
		case "slug":
			out.Slug = string(in.String())
		case "threads":
			out.Threads = int(in.Int())
		case "title":
			out.Title = string(in.String())
		case "user":
			out.Author = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeAForumInternalModels12(out *jwriter.Writer, in Forum) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"posts\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Posts))
	}
	{
		const prefix string = ",\"slug\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Slug))
	}
	{
		const prefix string = ",\"threads\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Threads))
	}
	{
		const prefix string = ",\"title\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"user\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Author))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Forum) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeAForumInternalModels12(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Forum) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeAForumInternalModels12(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Forum) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeAForumInternalModels12(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Forum) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeAForumInternalModels12(l, v)
}
func easyjsonD2b7633eDecodeAForumInternalModels13(in *jlexer.Lexer, out *Error) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "message":
			out.Message = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeAForumInternalModels13(out *jwriter.Writer, in Error) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Error) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeAForumInternalModels13(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Error) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeAForumInternalModels13(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Error) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeAForumInternalModels13(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Error) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeAForumInternalModels13(l, v)
}
