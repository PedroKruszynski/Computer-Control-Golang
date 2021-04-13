package volume_control

import (
	"fmt"
	"log"
	"net/http"

	"github.com/itchyny/volume-go"
)

func ShowVolumeActual(r *http.Request) string {
	vol, err := volume.GetVolume()
	if err != nil {
		log.Printf("Não deu para pegar o volume: %+v", err)
	}

	fmt.Println(vol)

	err = volume.SetVolume(65)
	if err != nil {
		log.Fatalf("Não deu para pegar o volume: %+v", err)
	}

	return "Volume aumentado com sucesso"
}
