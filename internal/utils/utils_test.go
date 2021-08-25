package utils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-location-api/internal/location"
	"github.com/ozonva/ova-location-api/internal/utils"
	"time"
)

var _ = Describe("Utils", func() {
	Describe("String Utils", func() {
		Describe("Проверка функции фильтрации слайса по слайсу", func() {
			var (
				sourceSlice  []string
				resultSlice  []string
				excludeSLice []string
			)

			BeforeEach(func() {
				sourceSlice = []string{"a", "b", "c", "d"}
				excludeSLice = []string{"b", "d"}
			})

			JustBeforeEach(func() {
				resultSlice = utils.SliceFilter(sourceSlice, excludeSLice)
			})

			Context("Размер входного слайса равен 0", func() {
				BeforeEach(func() {
					sourceSlice = []string{}
				})
				It("Должен вернуться nil", func() {
					Expect(resultSlice).Should(BeNil())
				})
			})

			Context("Размер слайса с исключениями равен 0", func() {
				BeforeEach(func() {
					excludeSLice = []string{}
				})
				It("Должен вернуться исходный слайс", func() {
					Expect(resultSlice).Should(BeEquivalentTo(sourceSlice))
				})
			})

			Context("Базовый сценарий", func() {
				It("Должен вернуться отфильтрованный слайс", func() {
					Expect(resultSlice).Should(BeEquivalentTo([]string{"a", "c"}))
				})
			})
		})

		Describe("Проверка функции разбивки на слайсы", func() {
			var (
				sourceSlice []string
				resultSlice [][]string
				chunkSize   int
			)

			BeforeEach(func() {
				sourceSlice = []string{"a", "b", "c", "d"}
			})

			JustBeforeEach(func() {
				resultSlice = utils.SliceSplit(sourceSlice, chunkSize)
			})

			Context("Размер входного слайса равен 0", func() {
				BeforeEach(func() {
					chunkSize = 5
					sourceSlice = []string{}
				})
				It("Должен вернуться пустой слайс", func() {
					Expect(resultSlice).Should(BeEquivalentTo([][]string{}))
				})
			})

			Context("Размер чанка больше размера входного слайса", func() {
				BeforeEach(func() {
					chunkSize = 5
				})
				It("Должен вернуться один чанк размера равного входному слайсу", func() {
					Expect(resultSlice).Should(BeEquivalentTo([][]string{sourceSlice}))
				})
			})

			Context("Размер чанка равен размеру входного слайса", func() {
				BeforeEach(func() {
					chunkSize = 4
				})
				It("Должен вернуться один чанк размера равного входному слайсу", func() {
					Expect(resultSlice).Should(BeEquivalentTo([][]string{sourceSlice}))
				})
			})

			Context("Размер чанка больше размера входного слайса", func() {
				BeforeEach(func() {
					chunkSize = 3
				})
				It("Должно вернуться два чанка общим размером равным входному слайсу", func() {
					Expect(resultSlice).Should(BeEquivalentTo([][]string{sourceSlice[0:3], sourceSlice[3:]}))
				})
			})

			Context("Размер чанка равен 0", func() {
				BeforeEach(func() {
					chunkSize = 0
				})
				It("Должен вернуться nil", func() {
					Expect(resultSlice).Should(BeNil())
				})
			})
		})

		Describe("Проверка функции перестановки ключей и значений местами", func() {
			Context("Базовый сценарий", func() {
				It("Должна вернуться мапа исходного размера, ключи и значения поменяны местами", func() {
					resultMap := utils.MapFlip(map[string]string{
						"BCN": "Barcelona",
						"LED": "Saint-Petersburg",
						"MSK": "Moscow",
					})
					Expect(resultMap).Should(BeEquivalentTo(map[string]string{
						"Barcelona":        "BCN",
						"Saint-Petersburg": "LED",
						"Moscow":           "MSK",
					}))
				})
			})

			Context("Есть коллизии", func() {
				It("Должна вернуться мапа меньшего размера, коллизии исключены, ключи и значения поменяны местами", func() {
					resultMap := utils.MapFlip(map[string]string{
						"BCN":  "Barcelona",
						"LED":  "Saint-Petersburg",
						"MSK2": "Moscow",
						"MSK":  "Moscow",
					})
					Expect(resultMap).Should(BeEquivalentTo(map[string]string{
						"Barcelona":        "BCN",
						"Saint-Petersburg": "LED",
						"Moscow":           "MSK2",
					}))
				})
			})

			Context("Размер входной мапы равен 0", func() {
				It("Должна вернуться пустая мапа", func() {
					resultMap := utils.MapFlip(map[string]string{})
					Expect(resultMap).Should(BeEquivalentTo(map[string]string{}))
				})
			})
		})
	})

	Describe("Location Utils", func() {
		Describe("Проверка функции разбивки на слайсы", func() {
			var (
				sourceSlice []location.Location
				resultSlice [][]location.Location
				chunkSize   int
			)

			BeforeEach(func() {
				sourceSlice = []location.Location{
					{1, 1, "Some address 1", 20, 30, time.Now()},
					{2, 2, "Some address 2", 22, 38, time.Now()},
					{3, 3, "Some address 3", 19, 31, time.Now()},
					{4, 4, "Some address 4", 17, 24, time.Now()},
				}
			})

			JustBeforeEach(func() {
				resultSlice = utils.LocationSliceSplit(sourceSlice, chunkSize)
			})

			Context("Размер входного слайса равен 0", func() {
				BeforeEach(func() {
					chunkSize = 5
					sourceSlice = []location.Location{}
				})
				It("Должен вернуться пустой слайс", func() {
					Expect(resultSlice).Should(BeEquivalentTo([][]location.Location{}))
				})
			})

			Context("Размер чанка больше размера входного слайса", func() {
				BeforeEach(func() {
					chunkSize = 5
				})
				It("Должен вернуться один чанк размера равного входному слайсу", func() {
					Expect(resultSlice).Should(BeEquivalentTo([][]location.Location{sourceSlice}))
				})
			})

			Context("Размер чанка равен размеру входного слайса", func() {
				BeforeEach(func() {
					chunkSize = 4
				})
				It("Должен вернуться один чанк размера равного входному слайсу", func() {
					Expect(resultSlice).Should(BeEquivalentTo([][]location.Location{sourceSlice}))
				})
			})

			Context("Размер чанка больше размера входного слайса", func() {
				BeforeEach(func() {
					chunkSize = 3
				})
				It("Должно вернуться два чанка общим размером равным входному слайсу", func() {
					Expect(resultSlice).Should(BeEquivalentTo([][]location.Location{sourceSlice[0:3], sourceSlice[3:]}))
				})
			})

			Context("Размер чанка равен 0", func() {
				BeforeEach(func() {
					chunkSize = 0
				})
				It("Должен вернуться nil", func() {
					Expect(resultSlice).Should(BeNil())
				})
			})
		})

		Describe("Проверка функции преобразования слайса локаций в мапу по Id локации", func() {
			Context("Базовый сценарий", func() {
				It("Должна вернуться мапа с Id в качестве ключа, размер соответствует размеру входного слайса", func() {
					time := time.Now()
					resultMap := utils.LocationSliceToMap([]location.Location{
						{1, 1, "Some address 1", 20, 30, time},
						{2, 2, "Some address 2", 22, 34, time},
					})
					Expect(resultMap).Should(BeEquivalentTo(map[uint64]location.Location{
						1: {1, 1, "Some address 1", 20, 30, time},
						2: {2, 2, "Some address 2", 22, 34, time},
					}))
				})
			})

			Context("Есть коллизии", func() {
				It("Должна вернуться мапа с Id в качестве ключа, коллизии по Id исключены", func() {
					time := time.Now()
					resultMap := utils.LocationSliceToMap([]location.Location{
						{1, 1, "Some address 1", 20, 30, time},
						{2, 2, "Some address 2", 22, 34, time},
						{2, 2, "Some address 2", 22, 34, time},
					})
					Expect(resultMap).Should(BeEquivalentTo(map[uint64]location.Location{
						1: {1, 1, "Some address 1", 20, 30, time},
						2: {2, 2, "Some address 2", 22, 34, time},
					}))
				})
			})

			Context("Размер входной мапы равен 0", func() {
				It("Должна вернуться пустая мапа", func() {
					resultMap := utils.LocationSliceToMap([]location.Location{})
					Expect(resultMap).Should(BeEquivalentTo(map[uint64]location.Location{}))
				})
			})
		})
	})
})
