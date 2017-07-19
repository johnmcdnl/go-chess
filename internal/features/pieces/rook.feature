Feature: Rook

  Background:
    Given an empty board

  Scenario Outline: A Rook in the middle of an empty board
    When a "black" "rook" is placed on the square "D4"
    And we print the board
    Then the piece has moves available to "<squares>"
    Examples:
      | squares |
      |         |