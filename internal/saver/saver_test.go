package saver

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-location-api/internal/location"
	"github.com/ozonva/ova-location-api/internal/mocks"
	"time"
)

var _ = Describe("Saver", func() {
	var (
		ctrl            *gomock.Controller
		mockFlusher     *mocks.MockFlusher
		locations       []location.Location
		capacity        uint64
		timeout         time.Duration
		saverInstance   *saver
		flushedEntities []location.Location
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockFlusher = mocks.NewMockFlusher(ctrl)
		locations = []location.Location{
			{1, 1, "Some address 1", 20, 30, time.Now()},
			{2, 2, "Some address 2", 22, 38, time.Now()},
			{3, 3, "Some address 3", 19, 31, time.Now()},
			{4, 4, "Some address 4", 17, 24, time.Now()},
		}
		capacity = 2
		timeout = 1 * time.Second
	})

	JustBeforeEach(func() {
		flushedEntities = make([]location.Location, 0)
		saverInstance = New(capacity, mockFlusher, timeout)
		saverInstance.Init()
	})

	Describe("Добавляем единоразово количество элементов меньшее, чем ёмкость буфера", func() {
		BeforeEach(func() {
			mockFlusher.EXPECT().Flush(gomock.Len(1)).DoAndReturn(func(entity []location.Location) []location.Location {
				flushedEntities = append(flushedEntities, entity...)
				return nil
			}).AnyTimes()
		})

		JustBeforeEach(func() {
			saverInstance.Save(locations[0])
		})

		Context("Проверяем добавление элементов", func() {
			It("Сразу после добавления в буфере один элемент", func() {
				Eventually(func() []location.Location {
					return saverInstance.buffer
				}, 100*time.Millisecond).Should(BeEquivalentTo([]location.Location{locations[0]}))
				saverInstance.Close()
			})
		})

		Context("Проверяем сброс элементов в хранилище", func() {
			It("По истечении таймаута буфер пуст", func() {
				Eventually(func() []location.Location {
					return saverInstance.buffer
				}, timeout+500*time.Millisecond).Should(BeNil())
				saverInstance.Close()
			})

			It("После вызова метода Close буфер пуст", func() {
				saverInstance.Close()
				Eventually(func() []location.Location {
					return saverInstance.buffer
				}, 100*time.Millisecond).Should(BeNil())
			})
		})
	})

	AfterEach(func() {
		ctrl.Finish()
	})
})
