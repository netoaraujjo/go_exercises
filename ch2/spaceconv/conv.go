package spaceconv

// FToM converte uma distância em pés par metros.
func FToM(f Foot) Meters { return Meters(f * 0.3048) }

// MToF converte uma distância em metros para pés.
func MToF(m Meters) Foot { return Foot(m * 3.28084) }
