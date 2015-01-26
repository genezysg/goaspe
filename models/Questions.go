package models



var Questionrepository *QuestionRepo


func init(){
Questionrepository=new(QuestionRepo)
Questionrepository.CreateQuestions()

}

type Question struct{
	Enun string
	Alt	[]string
}


type QuestionRepo struct{
	questions []*Question
}

func (this *QuestionRepo) CreateQuestions()	{
	q:=new(Question)
	q.Enun="Questao 1"
	q.Alt=append(q.Alt,"a")
	q.Alt=append(q.Alt,"b")
	q.Alt=append(q.Alt,"c")
	this.questions=append(this.questions,q)
	q=new(Question)
	q.Enun="Questao 2"
	q.Alt=append(q.Alt,"d")
	q.Alt=append(q.Alt,"e")
	q.Alt=append(q.Alt,"f")
	this.questions=append(this.questions,q)
}

func (this *QuestionRepo) Save(q *Question) {
	this.questions=append(this.questions,q)
}

func (this *QuestionRepo) GetAll() []*Question{
	return this.questions
}

func (this *QuestionRepo) FindById(id int64) (*Question, bool){
	for i,v:= range this.questions{
		if intid:=int(id); (i+1)==intid{
				return v,true
		}	
	}
				return nil,false
}
