package handlers

import (
    "$GO_BOOTSTRAP_REPO_NAME/$GO_BOOTSTRAP_REPO_USER/$GO_BOOTSTRAP_PROJECT_NAME/models"
    "$GO_BOOTSTRAP_REPO_NAME/$GO_BOOTSTRAP_REPO_USER/$GO_BOOTSTRAP_PROJECT_NAME/libhttp"
    "github.com/gorilla/context"
    "github.com/gorilla/sessions"
    "html/template"
    "net/http"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")

    cookieStore := context.Get(r, "cookieStore").(*sessions.CookieStore)

    session, _ := cookieStore.Get(r, "$GO_BOOTSTRAP_PROJECT_NAME-session")
    currentUser, ok := session.Values["user"].(*models.UserRow)
    if !ok {
        http.Redirect(w, r, "/logout", 302)
        return
    }

    data := struct {
        CurrentUser *models.UserRow
    }{
        currentUser,
    }

    tmpl, err := template.ParseFiles("templates/dashboard.html.tmpl", "templates/home.html.tmpl")
    if err != nil {
        libhttp.HandleErrorJson(w, err)
        return
    }

    tmpl.Execute(w, data)
}
