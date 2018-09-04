# grcir
Ranked Choice Instant Runoff voting library written in Go

In the result of a tie for last place and no winner, all the last place candidates are removed, so it is not a 'Spoiler Proof' method [Spoiler Effect + Resistance to tactical voting](https://en.wikipedia.org/wiki/Instant-runoff_voting#Resistance_to_tactical_voting) 

## Example
| Voter A | Voter B | Voter C | Voter D | Voter E |
|---------|---------|---------|---------|---------|
| Bob     | Sue     | Bill    | Bob     | Sue     |
| Bill    | Bob     | Sue     | Bill    | Bob     |
| Sue     | Bill    | Bob     | Sue     | Bill    |

This would result `[]string{ "Sue" }`

## Tie Example

| Voter A | Voter B |
|---------|---------|
| Bob     | Sue     |

This would result in an This would result `[]string{ "Sue", "Bob" }`
