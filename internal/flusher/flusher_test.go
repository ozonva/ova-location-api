package flusher_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-location-api/internal/flusher"
	"github.com/ozonva/ova-location-api/internal/location"
	"github.com/ozonva/ova-location-api/internal/mocks"
	"time"
)

var _ = Describe("Flusher", func() {
	var (
		ctrl            *gomock.Controller
		mockRepo        *mocks.MockRepo
		locations       []location.Location
		unsavedEntities []location.Location
		chunkSize       int
		flusherInstance flusher.Flusher
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
		locations = []location.Location{
			{1, 1, "Some address 1", 20, 30, time.Now()},
			{2, 2, "Some address 2", 22, 38, time.Now()},
			{3, 3, "Some address 3", 19, 31, time.Now()},
			{4, 4, "Some address 4", 17, 24, time.Now()},
		}
		chunkSize = 2
	})

	JustBeforeEach(func() {
		flusherInstance = flusher.NewFlusher(chunkSize, mockRepo)
		unsavedEntities = flusherInstance.Flush(locations)
	})

	Context("Передана некорректная величина чанка", func() {
		BeforeEach(func() {
			chunkSize = 0
		})
		It("Возвращаемое значение равно nil", func() {
			Expect(unsavedEntities).Should(BeNil())
		})
	})

	Context("Все чанки удалось сохранить", func() {
		BeforeEach(func() {
			mockRepo.EXPECT().AddEntities(gomock.Any()).Return(nil).AnyTimes()
		})
		It("Возвращаемое значение равно nil", func() {
			Expect(unsavedEntities).Should(BeNil())
		})
	})

	Context("Удалось сохранить только первый чанк", func() {
		BeforeEach(func() {
			mockRepo.EXPECT().AddEntities(gomock.Any()).Return(nil).Times(1)
			mockRepo.EXPECT().AddEntities(gomock.Any()).Return(errors.New("error")).AnyTimes()
		})
		It("Возвращаемое значение равно входному без первого чанка", func() {
			Expect(unsavedEntities).Should(BeEquivalentTo(locations[chunkSize:]))
		})
	})

	Context("Не удалось сохранить ни одного чанка", func() {
		BeforeEach(func() {
			mockRepo.EXPECT().AddEntities(gomock.Any()).Return(errors.New("error")).AnyTimes()
		})
		It("Возвращаемое значение равно входному", func() {
			Expect(unsavedEntities).Should(BeEquivalentTo(locations))
		})
	})

	AfterEach(func() {
		ctrl.Finish()
	})
})
