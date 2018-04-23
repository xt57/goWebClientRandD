//
//
//

func Upload(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Upload files\n")

    file, handler, err := r.FormFile("file")
    if err != nil {
        panic(err) //dont do this
    }
    defer file.Close()

    // copy example
    f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        panic(err) //please dont
    }
    defer f.Close()
    io.Copy(f, file)
}

