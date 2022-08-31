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