// Pacote weightconv realiza conversões de Kilogramas e Libras.
package weightconv

import "fmt"

type Kilo float64
type Libra float64

func (k Kilo) String() string   { return fmt.Sprintf("%gkg", k) }
func (l Libra) String() string { return fmt.Sprintf("%glb", l) }
