package clientGRPCJsonXml
type PhoneNumber struct {
	Number string `json:"number" xml:"number"`
	Kind   int    `json:"kind" xml:"kind"`
}
type Person struct {
	Name string `json:"name" xml:"name"`
	Id int `json:"id" xml:"id"`
	Email string `json:"email" xml:"email"`
	Phones []PhoneNumber `json:"phones" xml:"phones"`
}
type ResponseEdit struct {
	Result bool `json:"result" xml:"result"`
}
type RequestId struct {
	Id int32 `json:"id" xml:"id"`
}