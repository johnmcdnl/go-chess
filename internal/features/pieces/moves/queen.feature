@queen
Feature: Queen

  Background:
    Given an empty board

  Scenario Outline: Piece in a corner of the empty board
    When a "<color>" "<piece>" is placed on the square "<origin>"
    Then the piece has moves available to "<validMoves>"
    Examples:
      | piece | color | origin | validMoves                                                           |
      | queen | black | A1     | A2 A3 A4 A5 A6 A7 A8 B1 C1 D1 E1 F1 G1 H1 B2 C3 D4 E5 F6 G7 H8       |
      | queen | black | H8     | H7 H6 H5 H4 H3 H2 H1 G8 F8 E8 D8 C8 B8 A8 G7 F6 E5 D4 C3 B2 A1       |
      | queen | black | B7     | A7 C7 D7 E7 F7 G7 H7 B8 B6 B5 B4 B3 B2 B1 A6 C8 A8 C6 D5 E4 F3 G2 H1 |

  Scenario Outline: Piece has valid capture moves
    When a "<oppColor>" "<oppPiece>" is placed on the square "<oppSquare>"
    When a "<color>" "<piece>" is placed on the square "<origin>"
    Then the piece has moves available to "<validMoves>"
    Examples:
      | piece | color | origin | oppPiece | oppColor | oppSquare | validMoves                                               |
      | queen | white | A1     | rook     | black    | B2        | A2 A3 A4 A5 A6 A7 A8 B1 C1 D1 E1 F1 G1 H1 B2             |
      | queen | white | B7     | rook     | black    | D5        | A7 C7 D7 E7 F7 G7 H7 B8 B6 B5 B4 B3 B2 B1 A6 C8 A8 C6 D5 |
