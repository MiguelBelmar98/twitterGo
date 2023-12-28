package routers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/miguelbelmar98/twittergo/models"
	"github.com/miguelbelmar98/twittergo/bd"
)

func Registro(ctx context.Context) models.RespApi {
	var t models.Usuario
	var r models.RespApi

	r.StatusCode = 400

	fmt.Println("Entre a Registro")

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		fmt.Println(r.Message)
		r.Message = err.Error()
		return r
	}
	if len(t.Email) == 0 {
		r.Message = "Email requerido"
		fmt.Println(r.Message)
		return r
	}
	if len(t.Password) < 6 {
		r.Message = "Password requerido y mayor a 6 caracteres"
		fmt.Println(r.Message)
		return r
	}
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado {	
		r.Message = "Ya existe un usuario registrado con ese email"
		fmt.Println(r.Message)
		return r
	}
	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		r.Message = "Ocurrio un error al intentar realizar el registro de usuario" + err.Error()
		fmt.Println(r.Message)
		return r
	}
	if !status {
		r.Message = "No se ha logrado insertar el registro de usuario"
		fmt.Println(r.Message)
		return r
	}
	r.StatusCode = 200
	r.Message = "Registro exitoso"
	fmt.Println(r.Message)
	return r
}
