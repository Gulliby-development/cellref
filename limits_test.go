package main

import "testing"

var rowsLimits = []row{
	{[]string{"check"}, "n0 A1\nn1 B1\nn2 C1\nn3 D1\nn4 E1\nn5 F1\nn6 G1\nn7 H1\nn8 I1\nn9 J1\nn10 A2\nn11 B2\nn12 C2\nn13 D2\nn14 E2\nn15 F2\nn16 G2\nn17 H2\nn18 I2\nn19 J2\nn20 A3\nn21 B3\nn22 C3\nn23 D3\nn24 E3\nn25 F3\nn26 G3\nn27 H3\nn28 I3\nn29 J3\nn30 A4\nn31 B4\nn32 C4\nn33 D4\nn34 E4\nn35 F4\nn36 G4\nn37 H4\nn38 I4\nn39 J4\n", 0, "ok 40 ranges\n", ""},
	{[]string{"check"}, "n0 A1\nn1 B1\nn2 C1\nn3 D1\nn4 E1\nn5 F1\nn6 G1\nn7 H1\nn8 I1\nn9 J1\nn10 A2\nn11 B2\nn12 C2\nn13 D2\nn14 E2\nn15 F2\nn16 G2\nn17 H2\nn18 I2\nn19 J2\nn20 A3\nn21 B3\nn22 C3\nn23 D3\nn24 E3\nn25 F3\nn26 G3\nn27 H3\nn28 I3\nn29 J3\nn30 A4\nn31 B4\nn32 C4\nn33 D4\nn34 E4\nn35 F4\nn36 G4\nn37 H4\nn38 I4\nn39 J4\nn40 A5\n", 1, "", "cellref: input is too long (max 40 lines)\n"},
	{[]string{"size"}, "n0 A1\nn1 B1\nn2 C1\nn3 D1\nn4 E1\nn5 F1\nn6 G1\nn7 H1\nn8 I1\nn9 J1\nn10 A2\nn11 B2\nn12 C2\nn13 D2\nn14 E2\nn15 F2\nn16 G2\nn17 H2\nn18 I2\nn19 J2\nn20 A3\nn21 B3\nn22 C3\nn23 D3\nn24 E3\nn25 F3\nn26 G3\nn27 H3\nn28 I3\nn29 J3\nn30 A4\nn31 B4\nn32 C4\nn33 D4\nn34 E4\nn35 F4\nn36 G4\nn37 H4\nn38 I4\nn39 J4\n", 0, "n0 A1 1x1 1 cell\nn1 B1 1x1 1 cell\nn2 C1 1x1 1 cell\nn3 D1 1x1 1 cell\nn4 E1 1x1 1 cell\nn5 F1 1x1 1 cell\nn6 G1 1x1 1 cell\nn7 H1 1x1 1 cell\nn8 I1 1x1 1 cell\nn9 J1 1x1 1 cell\nn10 A2 1x1 1 cell\nn11 B2 1x1 1 cell\nn12 C2 1x1 1 cell\nn13 D2 1x1 1 cell\nn14 E2 1x1 1 cell\nn15 F2 1x1 1 cell\nn16 G2 1x1 1 cell\nn17 H2 1x1 1 cell\nn18 I2 1x1 1 cell\nn19 J2 1x1 1 cell\nn20 A3 1x1 1 cell\nn21 B3 1x1 1 cell\nn22 C3 1x1 1 cell\nn23 D3 1x1 1 cell\nn24 E3 1x1 1 cell\nn25 F3 1x1 1 cell\nn26 G3 1x1 1 cell\nn27 H3 1x1 1 cell\nn28 I3 1x1 1 cell\nn29 J3 1x1 1 cell\nn30 A4 1x1 1 cell\nn31 B4 1x1 1 cell\nn32 C4 1x1 1 cell\nn33 D4 1x1 1 cell\nn34 E4 1x1 1 cell\nn35 F4 1x1 1 cell\nn36 G4 1x1 1 cell\nn37 H4 1x1 1 cell\nn38 I4 1x1 1 cell\nn39 J4 1x1 1 cell\ntotal 40 cells\n", ""},
	{[]string{"check"}, "n0 A1\nn1 B1\nn2 C1\nn3 D1\nn4 E1\nn5 F1\nn6 G1\nn7 H1\nn8 I1\nn9 J1\nn10 A2\nn11 B2\nn12 C2\nn13 D2\nn14 E2\nn15 F2\nn16 G2\nn17 H2\nn18 I2\nn19 J2\nn20 A3\nn21 B3\nn22 C3\nn23 D3\nn24 E3\nn25 F3\nn26 G3\nn27 H3\nn28 I3\nn29 J3\nn30 A4\nn31 B4\nn32 C4\nn33 D4\nn34 E4\nn35 F4\nn36 G4\nn37 H4\nn38 I4\nn39 J4\nx", 1, "", "cellref: input is too long (max 40 lines)\n"},
}

func TestRowsLimits(t *testing.T) { checkRows(t, rowsLimits) }
