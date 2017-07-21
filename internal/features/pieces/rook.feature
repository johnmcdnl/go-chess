Feature: Rook

  Background:
    Given an empty board

  Scenario Outline: Piece in the middle of an empty board
    When a "<color>" "<piece>" is placed on the square "<origin>"
    Then the piece has moves available to "<validMoves>"
    Examples:
      | piece  | color | origin | validMoves                                |
#      | rook   | black | D4     | D1 D2 D3 D5 D6 D7 D8 A4 B4 C4 E4 F4 G4 H4 |
#      | knight | black | D4     | B5 C6 B3 C2 E2 E6 F3 F5                   |
#      | king   | black | D4     | C3 C4 C5 D3 D5 E3 E4 E5                   |

  Scenario Outline: Piece Placed in corner of the board
    When a "<color>" "<piece>" is placed on the square "<origin>"
    Then the piece has moves available to "<validMoves>"
    Examples:
      | piece  | color | origin | validMoves                                |
#      | rook   | white | A1     | A2 A3 A4 A5 A6 A7 A8 B1 C1 D1 E1 F1 G1 H1 |
#      | knight | white | A1     | B3 C2                                     |
#      | king   | white | A1     | A2 B2 B1                                  |

  Scenario Outline: Piece Placed in corner of the board
    Given an empty board
    When a "<oppColor>" "<oppPiece>" is placed on the square "<oppSquare>"
    When a "<color>" "<piece>" is placed on the square "<origin>"
    Then the piece has moves available to "<validMoves>"
    Examples:
      | piece  | color | origin | oppPiece | oppColor | oppSquare | validMoves                 |
      | rook   | white | A1     | rook     | black    | A3        | A2 A3 B1 C1 D1 E1 F1 G1 H1 |
      | knight | black | A1     | rook     | white    | B3        | B3 C2                      |
      | knight | black | A1     | rook     | white    | B2        | B3 C2                      |
      | king   | white | A1     | rook     | black    | A2        | A2 B2 B1                   |
      | king   | white | A1     | rook     | black    | A3        | A2 B2 B1                   |