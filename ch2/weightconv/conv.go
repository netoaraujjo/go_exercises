package weightconv

// FToM converte um peso em Kilogramas para Libras.
func KToL(k Kilo) Libra { return Libra(k * 2.20462) }

// MToF converte um peso em Libras para Kilogramas.
func LToK(l Libra) Kilo { return Kilo(l * 0.453592) }
