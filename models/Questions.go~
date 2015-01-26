package models




var QuestionRep *Question

func init(){
QuestionRep=new(Question)
}

type Question struct{
	Enun string
	Alt	[]string
}

func CreateQuestions() []*Question{
	var questions []*Question
	q:=new(Question)
	q.Enun="Questao 1"
	q.Alt=append(q.Alt,"a")
	q.Alt=append(q.Alt,"b")
	q.Alt=append(q.Alt,"c")
	questions=append(questions,q)
	q=new(Question)
	q.Enun="Questao 2"
	q.Alt=append(q.Alt,"d")
	q.Alt=append(q.Alt,"e")
	q.Alt=append(q.Alt,"f")
	questions=append(questions,q)
	return questions

}

func (this *Question) GetAll() []*Question{
	return CreateQuestions()
}

func (this *Question) FindById(id int64) (*Question, bool){
	questions:=CreateQuestions()	
	
	for i,v:= range questions{
		if intid:=int(id); i==intid{
				return v,true
		}	
	}
				return nil,false
}
