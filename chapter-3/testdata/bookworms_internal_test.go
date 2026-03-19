package main

import (
    "reflect"
    "testing"
)

var (
    handmaidsTale = Book{
        Author: "Margaret Atwood", Title: "The Handmaid's Tale",
    }
    oryxAndCrake  = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
    theBellJar    = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
    janeEyre      = Book{Author: "Charlotte Brontë", Title: "Jane Eyre"}
)

func TestLoadBookworms_Success(t *testing.T) {
    tests := map[string]struct {
        bookwormsFile string
        want          []Bookworm
        wantErr       bool
    }{
        "file exists": {
            bookwormsFile: "testdata/bookworms.json",
            want: []Bookworm{
                {Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
                {Name: "Peggy", Books: 
                    []Book{oryxAndCrake, handmaidsTale, janeEyre},
                },
            },
            wantErr: false,
        },
        "file doesn't exist": {...},
        "invalid JSON": {...},
    }
    for name, testCase := range tests {
        t.Run(name, func(t *testing.T) {
            got, err := loadBookworms(testCase.bookwormsFile)
            if err != nil && !testCase.wantErr { #1
                t.Fatalf("expected no error %s, got none", err.Error())
            }

            if err == nil && testCase.wantErr { #2
                t.Fatalf("expected an error, got one %s", err.Error())
            }

            if !equalBookworms(got, testCase.want) { #3
                t.Fatalf("different result: got %v, expected %v", 
                                    got, testCase.want)
            }
        })
    }
}

// // type hinting by giving the struct a name
// type testCase struct {
//     bookwormsFile string
//     want          []Bookworm
//     wantErr       bool
// }

// tests := map[string]testCase{
//     "file exists": testCase{
//         bookwormsFile: "testdata/bookworms.json",
//         want: []Bookworm{
//             {Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
//             {Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
//         },
//         wantErr: false,
//     },
// 		"name of test case" : testCase {
//       code within the test case
//		}
// }
