package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type (
	Siswa struct {
		NIS    int    `json:"nis"`
		Nama   string `json:"nama"`
		Kelas  int    `json:"kelas"`
		Gender string `json:"gender"`
	}

	Guru struct {
		NIG    int    `json:"nig"`
		Nama   string `json:"nama"`
		Gender string `json:"gender"`
	}

	MataPelajaran struct {
		ID    int    `json:"id"`
		Nama  string `json:"nama"`
		NIG   int    `json:"nig"`
		Kelas int    `json:"kelas"`
	}

	MataPelajaranByNis struct {
		NIS           int      `json:"nis"`
		MataPelajaran []string `json:"mata_pelajaran"`
	}
)

var students = []Siswa{
	{
		NIS:    1,
		Nama:   "Azmi",
		Kelas:  4,
		Gender: "L",
	},
	{
		NIS:    2,
		Nama:   "Dhaby",
		Kelas:  5,
		Gender: "L",
	},
	{
		NIS:    3,
		Nama:   "Fitria",
		Kelas:  6,
		Gender: "P",
	},
}

var Teachers = []Guru{
	{
		NIG:    1,
		Nama:   "Guru Pertama",
		Gender: "L",
	},
	{
		NIG:    2,
		Nama:   "Guru Kedua",
		Gender: "P",
	},
	{
		NIG:    3,
		Nama:   "Guru Ketiga",
		Gender: "P",
	},
}

var Lessons = []MataPelajaran{
	{
		ID:    1,
		Nama:  "IPA",
		NIG:   1,
		Kelas: 4,
	},
	{
		ID:    2,
		Nama:  "IPA",
		NIG:   2,
		Kelas: 5,
	},
	{
		ID:    3,
		Nama:  "IPA",
		NIG:   3,
		Kelas: 6,
	},
	{
		ID:    4,
		Nama:  "IPS",
		NIG:   1,
		Kelas: 4,
	},
	{
		ID:    5,
		Nama:  "MTK",
		NIG:   2,
		Kelas: 5,
	},
}

func Student() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			// get all data siswa
			response.WriteHeader(200)
			dataSiswaJson, err := json.Marshal(students)
			if err != nil {
				panic(err)
			}
			_, err = response.Write(dataSiswaJson)
			if err != nil {
				panic(err)
			}
			return
		} else if request.Method == http.MethodPost {
			// get 1 data siswa by nis
			// ambil data siswa dari body
			defer request.Body.Close()
			dataBody, err := io.ReadAll(request.Body)
			if err != nil {
				panic(err)
			}
			nis := struct {
				NIS int `json:"nis"`
			}{}
			err = json.Unmarshal(dataBody, &nis)
			if err != nil {
				panic(err)
			}

			dataSiswa := Siswa{NIS: 0}
			for _, val := range students {
				if val.NIS == nis.NIS {
					dataSiswa = val
				}
			}
			if dataSiswa.NIS == 0 {
				response.WriteHeader(404)
				_, err := response.Write([]byte("Data not found"))
				if err != nil {
					panic(err)
				}
				return
			}
			dataSiswaJson, err := json.Marshal(dataSiswa)
			if err != nil {
				panic(err)
			}
			_, err = response.Write(dataSiswaJson)
			if err != nil {
				panic(err)
			}
			return
		}

		response.WriteHeader(405)
		_, err := response.Write([]byte("Method is not allowed"))
		if err != nil {
			panic(err)
		}
		return
	}
}

func Teacher() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			// get all data guru
			response.WriteHeader(200)
			dataGuruJson, err := json.Marshal(Teachers)
			if err != nil {
				panic(err)
			}
			_, err = response.Write(dataGuruJson)
			if err != nil {
				panic(err)
			}
			return
		} else if request.Method == http.MethodPost {
			// get 1 data guru by nig
			// ambil data guru dari body
			defer request.Body.Close()
			dataBody, err := io.ReadAll(request.Body)
			if err != nil {
				panic(err)
			}
			nig := struct {
				NIG int `json:"nig"`
			}{}
			err = json.Unmarshal(dataBody, &nig)
			if err != nil {
				panic(err)
			}

			dataGuru := Guru{NIG: 0}
			for _, val := range Teachers {
				if val.NIG == nig.NIG {
					dataGuru = val
				}
			}
			if dataGuru.NIG == 0 {
				response.WriteHeader(404)
				_, err := response.Write([]byte("Data not found"))
				if err != nil {
					panic(err)
				}
				return
			}
			dataGuruJson, err := json.Marshal(dataGuru)
			if err != nil {
				panic(err)
			}
			_, err = response.Write(dataGuruJson)
			if err != nil {
				panic(err)
			}
			return
		}

		response.WriteHeader(405)
		_, err := response.Write([]byte("Method is not allowed"))
		if err != nil {
			panic(err)
		}
		return
	}
}

func Lesson() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			// get 1 data mapel by id
			// ambil data mapel dari body
			defer request.Body.Close()
			dataBody, err := io.ReadAll(request.Body)
			if err != nil {
				panic(err)
			}
			id := struct {
				ID int `json:"id"`
			}{}
			err = json.Unmarshal(dataBody, &id)
			if err != nil {
				panic(err)
			}

			dataMapel := MataPelajaran{ID: 0}
			for _, val := range Lessons {
				if val.ID == id.ID {
					dataMapel = val
				}
			}
			if dataMapel.ID == 0 {
				response.WriteHeader(404)
				_, err := response.Write([]byte("Data not found"))
				if err != nil {
					panic(err)
				}
				return
			}
			dataMapelJson, err := json.Marshal(dataMapel)
			if err != nil {
				panic(err)
			}
			_, err = response.Write(dataMapelJson)
			if err != nil {
				panic(err)
			}
			return
		}
		response.WriteHeader(405)
		_, err := response.Write([]byte("Method is not allowed"))
		if err != nil {
			panic(err)
		}
		return
	}
}

func LessonByNis() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			nis := request.URL.Query().Get("nis")
			if nis == "" {
				response.WriteHeader(400)
				_, err := response.Write([]byte("Memerlukan NIS untuk mengetahui mata pelajaran"))
				if err != nil {
					panic(err)
				}
				return
			}

			nisInt, err := strconv.Atoi(nis)
			if err != nil {
				response.WriteHeader(400)
				_, err := response.Write([]byte("Nis tidak ditemukan"))
				if err != nil {
					panic(err)
				}
				return
			}

			var mataPelajaran []string
			for _, student := range students {
				if student.NIS == nisInt {
					for _, lesson := range Lessons {
						if lesson.Kelas == student.Kelas {
							mataPelajaran = append(mataPelajaran, lesson.Nama)
						}
					}
				}
			}

			fmt.Println("Mata Pelajaran:", mataPelajaran)
			data := MataPelajaranByNis{
				NIS:           nisInt,
				MataPelajaran: mataPelajaran,
			}
			response.WriteHeader(http.StatusOK)
			err = json.NewEncoder(response).Encode(data)
			if err != nil {
				panic(err)
			}
			return
		}
	}
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	// routing
	mux.HandleFunc("/students", Student())
	mux.HandleFunc("/teachers", Teacher())
	mux.HandleFunc("/lesson", Lesson())
	mux.HandleFunc("/lessonByNIS", LessonByNis())
	return mux
}

func main() {
	http.ListenAndServe("localhost:8080", GetMux())
}
