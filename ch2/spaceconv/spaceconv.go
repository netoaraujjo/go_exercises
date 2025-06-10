// Pacote tempconv realiza conversões de Pés e Metros.
package spaceconv

import "fmt"

type Foot float64
type Meters float64

func (f Foot) String() string   { return fmt.Sprintf("%g ft", f) }
func (m Meters) String() string { return fmt.Sprintf("%g m", m) }
