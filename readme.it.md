# gophysics
![GitHub](https://img.shields.io/github/license/Gabri432/gophysics)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Gabri432/gophysics)
[![Go Reference](https://pkg.go.dev/badge/github.com/Gabri432/gophysics.svg)](https://pkg.go.dev/github.com/Gabri432/gophysics)

Una collezione di alcune delle formule e costanti più comuni in fisica.

## Funzionalità
- Fornisce formule e costanti fisiche per fare calcoli.
- Permette la creazione di alcuni oggetti per semplificare certi calcoli (vedi esempio).
- Fornisce anche formule matematiche comuni alla fisica.

## Quali formule sono contenute?
### Meccanica Classica
- Forza, Velocità, Tempo, Lavoro, Accelerazione, Densità, Intensità, Potenza, Momento
- Energia Potenziale, Meccanica, Cinetica, Frequenza, Effetto Doppler.
### Gravità
- Energia potenziale Gravitazionale, Legge di Gravitazione Universale, Campo gravitazionale
- Velocità di fuga

### Fluidi
- Legge di Stokes, legge di Hagen-Poiseuille

### Elettromagnetismo
- Legge di Ohm, Capacità di un Conduttore, legge di Coulomb, Densità di energia.
- Campo Elettrico, Differenza di Energia Potenziale Elettrica.

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
```go
package main

import (
    "fmt"
    "github.com/Gabri432/gophysics/constants"
    "github.com/Gabri432/gophysics/formulas"
)

func main() {
    fmt.Println(constants.C)      <<< Scrivi una costante
    fmt.Println(formulas.Force(3, 4))   <<< Chiama una funzione
    myPlanet := formulas.PlanetBody{Mass: constants.EARTH_MASS, Radius: constants.EARTH_RADIUS}  <<< Costrusci il tuo oggetto 
	fmt.Println(myPlanet.EscapeSpeed())    <<< Chiama i suoi metodi per fare calcoli più comodamente.
}

=== Output ===
299792453  
12 N      <<< Quando si chiamano le funzioni queste ritornano un valore e la unità di misura.
11183.719071923773 m/s

```

- Using mathematical functions
```go
import (
    "fmt"
    "github.com/Gabri432/gophysics/mathem"
)

func main() {
    fmt.Println(mathem.Pi) <<< Writing a constant
}

=== Output ===
3.1415926535

```

## Struttura del progetto
### Cartelle

#### gophysics (main)
- `gophysics.go`, non fornisce alcuna formula, ma diverse funzioni per semplificare alcune operazioni.
- `license`, `readme.md`, `readme.it.md`, `CHANGELOG.txt`.

#### formulas
- `classical.go`, dove tutte le formule di FIsica Classica sono presenti.
- `fluids.go`, dove tutte le formule sui Fluidi sono presenti.
- `gravity.go`, dove tutte le formule sulla Gravità sono presenti.
- `thermodynamics.go`, dove tutte le formule sulla Termodinamica sono presenti.
- `electromagnetism.go`, dove tutte le formule sull'Elettromagnetismo sono presenti.
- `relativity.go`, where all the Relativity formulas are located.

#### constants
- `constants.go`, dove tutte le formule sulla Relatività sono presenti.

#### mathem 
- `constants.go`, dove tutte le costanti matematiche sono presenti.
- `conversions.go`, dove tutte le principali funzioni di calcolo matematico sono presenti.

## Contribuire al progetto
- Se vuoi aggiungere una feature o fare un aggiustamento fai un giro in questa pagina su come fare: [Contribuire a gophysics](https://github.com/Gabri432/gophysics/blob/master/.github/CONTRIBUTING.it.md)

## Nota bene
- Le formule sono state prese dal seguente libro : Titolo - "Fisica volume 1", Autori - ["Paolo Mazzoldi", "Massimo Nigro", "Cesare Voci"].