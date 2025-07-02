package service

import (
	"Ainotes/dto"
	"Ainotes/model"
	"Ainotes/repository"

	"github.com/google/uuid"
)

type QuizServiceImpl struct {
	quizRepo repository.QuizRepository
}

func NewQuizService(quizRepo repository.QuizRepository) QuizService {
	return &QuizServiceImpl{quizRepo: quizRepo}
}

func (svc *QuizServiceImpl) Create(req dto.QuizCreateRequest) (dto.QuizResponse, error) {
	quiz := model.Quiz{
		ID:                uuid.New(),
		NoteID:            req.NoteID,
		Question:          req.Question,
		Options:           req.Options,
		CorrectAnswerIndex: req.CorrectAnswerIndex,
		Explanation:       req.Explanation,
	}
	if err := svc.quizRepo.Create(&quiz); err != nil {
		return dto.QuizResponse{}, err
	}
	return dto.QuizResponse{
		ID:                quiz.ID,
		NoteID:            quiz.NoteID,
		Question:          quiz.Question,
		Options:           quiz.Options,
		CorrectAnswerIndex: quiz.CorrectAnswerIndex,
		Explanation:       quiz.Explanation,
	}, nil
}

func (svc *QuizServiceImpl) FindByID(id uuid.UUID) (dto.QuizResponse, error) {
	quiz, err := svc.quizRepo.FindByID(id)
	if err != nil {
		return dto.QuizResponse{}, err
	}
	return dto.QuizResponse{
		ID:                quiz.ID,
		NoteID:            quiz.NoteID,
		Question:          quiz.Question,
		Options:           quiz.Options,
		CorrectAnswerIndex: quiz.CorrectAnswerIndex,
		Explanation:       quiz.Explanation,
	}, nil
}

func (svc *QuizServiceImpl) FindAll() ([]dto.QuizResponse, error) {
	quizzes, err := svc.quizRepo.FindAll()
	if err != nil {
		return nil, err
	}
	responses := make([]dto.QuizResponse, len(quizzes))
	for i, quiz := range quizzes {
		responses[i] = dto.QuizResponse{
			ID:                quiz.ID,
			NoteID:            quiz.NoteID,
			Question:          quiz.Question,
			Options:           quiz.Options,
			CorrectAnswerIndex: quiz.CorrectAnswerIndex,
			Explanation:       quiz.Explanation,
		}
	}
	return responses, nil
}

func (svc *QuizServiceImpl) Update(id uuid.UUID, req dto.QuizUpdateRequest) (dto.QuizResponse, error) {
	quiz, err := svc.quizRepo.FindByID(id)
	if err != nil {
		return dto.QuizResponse{}, err
	}
	quiz.Question = req.Question
	quiz.Options = req.Options
	quiz.CorrectAnswerIndex = req.CorrectAnswerIndex
	quiz.Explanation = req.Explanation
	if err := svc.quizRepo.Update(quiz); err != nil {
		return dto.QuizResponse{}, err
	}
	return dto.QuizResponse{
		ID:                quiz.ID,
		NoteID:            quiz.NoteID,
		Question:          quiz.Question,
		Options:           quiz.Options,
		CorrectAnswerIndex: quiz.CorrectAnswerIndex,
		Explanation:       quiz.Explanation,
	}, nil
}

func (svc *QuizServiceImpl) Delete(id uuid.UUID) error {
	return svc.quizRepo.Delete(id)
}