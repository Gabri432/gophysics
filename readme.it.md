# gophysics
Una collezione di alcune delle formule e costanti più comuni in fisica.

## Quali formule sono contenute?
### Meccanica Classica
- Forza, Velocità, Tempo, Lavoro, Accelerazione, Densità, Intensità, Potenza, Momento
- Energia Potenziale, Meccanica, Cinetica
### Gravità
- Energia potenziale Gravitazionale, Legge di Gravitazione Universale, Campo gravitazionale
- Velocità di fuga

### Fluidi
- Legge di Stokes, legge di Hagen-Poiseuille

### Elettromagnetismo
- Legge di Ohm, Capacità di un Conduttore, legge di Coulomb, Effetto Doppler, Densità di energia.

### Termodinamica
- Legge di Gay-Lussac, Energia di calore netta, Flusso di calore, FLusso del campo elettrico, Calore di Joule

### Relatività
- Fattore di Lorentz, Tempo, Distanza, Massa e Momento relativistici
- Effetto fotoelettrico, velocità di Deriva.


## Come si usa
- Scarica 
```
go get github.com/Gabri432/gophysics
```

- Esempio di utilizzo
```
package main

import (
    "fmt"
    "github.com/Gabri432/gophysics
)

func main() {
    fmt.Println(gophysics.C)
}

=== Output ===
299792453  

```


## Struttura del progetto
### File
- `gophysics.go`, non fornisce alcuna formula, ma diverse funzioni per semplificare alcune operazioni.
- `classical.go`, dove tutte le formule di FIsica Classica sono presenti.
- `fluids.go`, dove tutte le formule sui Fluidi sono presenti.
- `gravity.go`, dove tutte le formule sulla Gravità sono presenti.
- `thermodinamics.go`, dove tutte le formule sulla Termodinamica sono presenti.
- `electromagnetism.go`, dove tutte le formule sull'Elettromagnetismo sono presenti.
- `relativity.go`, dove tutte le formule sulla Relatività sono presenti.
- `constants.go`, dove tutte le costanti sono presenti.
- `license`, `readme.md`, `readme.it.md`, `CHANGELOG.txt`.

## Contribuire al progetto
- Se vuoi aggiungere una feature o fare un aggiustamento fai un giro in questa pagina su come fare: [Contribuire a gophysics](https://github.com/Gabri432/gophysics/blob/master/.github/CONTRIBUTING.it.md)

## Nota bene
- Le formule sono state prese dal seguente libro : Titolo - "Fisica volume 1", Autori - ["Paolo Mazzoldi", "Massimo Nigro", "Cesare Voci"].