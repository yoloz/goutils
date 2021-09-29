package control

import (
	"birthdayreminder/server"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type DeleteControl struct{}

func (dc *DeleteControl) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "text/html; charset=utf-8")
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("parse form data fail,%v\n", err)
		http.Redirect(rw, r, "#!/err", http.StatusFound)
	} else {
		id, _ := strconv.Atoi(r.FormValue("id"))
		// fmt.Printf("delete birthday: %v\n", id)
		err := server.DeleteDb(id)
		if err != nil {
			fmt.Printf("delete birthday fail,%v\n", err)
			http.Redirect(rw, r, "#!/err", http.StatusFound)
		} else {
			http.Redirect(rw, r, "#!/list", http.StatusFound)
		}
	}
}

type QueryControl struct{}

func (qc *QueryControl) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json; charset=utf-8")
	l := server.QueryAll()
	var builder strings.Builder
	for e := l.Front(); e != nil; e = e.Next() {
		birthday, _ := e.Value.(server.Birthday)
		jsb, _ := json.Marshal(birthday)
		builder.WriteString(string(jsb))
		builder.WriteString(",")
	}
	var resp string
	if l.Len() == 0 {
		resp = "{\"data\":[],\"offset\":0,\"limit\":0,\"total\":0}"
	} else {
		resp = "{\"offset\":0,\"limit\":" + strconv.Itoa(l.Len()) + ",\"total\":" + strconv.Itoa(l.Len()) + ",\"data\":[" +
			builder.String()[0:builder.Len()-1] + "]}"
	}
	// fmt.Printf("query birtyday: %v\n", resp)
	io.WriteString(rw, resp)
}

type UpdateControl struct{}

func (uc *UpdateControl) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "text/html; charset=utf-8")
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("parse form data fail,%v\n", err)
		http.Redirect(rw, r, "#!/err", http.StatusFound)
	} else {
		id, _ := strconv.Atoi(r.FormValue("id"))
		timeType, _ := strconv.Atoi(r.FormValue("timeType"))

		birthday := server.Birthday{
			Id:        id,
			Name:      r.FormValue("name"),
			TimeType:  timeType,
			TimeText:  r.FormValue("timeText"),
			SendEmail: r.FormValue("sendEmail"),
		}
		// fmt.Printf("update birthday: %v\n", birthday)
		err := server.UpdateDb(birthday)
		if err != nil {
			fmt.Printf("update birthday fail,%v\n", err)
			http.Redirect(rw, r, "#!/err", http.StatusFound)
		} else {
			http.Redirect(rw, r, "#!/list", http.StatusFound)
		}
	}
}

type AddControl struct{}

func (ac *AddControl) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "text/html; charset=utf-8")
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("parse form data fail,%v\n", err)
		http.Redirect(rw, r, "#!/err", http.StatusFound)
	} else {
		timeType, _ := strconv.Atoi(r.FormValue("timeType"))
		birthday := server.Birthday{
			Id:        0,
			Name:      r.FormValue("name"),
			TimeType:  timeType,
			TimeText:  r.FormValue("timeText"),
			SendEmail: r.FormValue("sendEmail"),
		}
		// fmt.Printf("add new birthday: %v\n", birthday)
		err := server.InsertDb(birthday)
		if err != nil {
			fmt.Printf("add birthday fail,%v\n", err)
			http.Redirect(rw, r, "#!/err", http.StatusFound)
		} else {
			http.Redirect(rw, r, "#!/list", http.StatusFound)
		}
	}
}
