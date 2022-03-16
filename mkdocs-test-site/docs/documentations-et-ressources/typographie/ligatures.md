# Ligatures
Les **ligatures** de la police [JetBrains Mono](https://www.jetbrains.com/lp/mono/) fonctionnent / **s’affichent** manifestement **de manière contextuelle**. Je n’en sais pas la raison, mais il doit y en avoir une (éventuellement *bonne*).

Par exemple dans le présent fichier :
1. Marchent en dehors de `code` :
  - ✅ ==
  - ✅ `==`
  - ✅ != 
  - ✅ `!=`
2. Marchent pas en dehors de `code` :
  - ❌ ->
  - ✅ `->`
  - ❌ =>
  - ✅ `=>`

—Maylis le 20210707

## Réalisation des ligatures
### Différent
Le caractère `!=`
s’écrit
`!` + `=`

### Double-égal (ou égal égal)
Le caractère `==`
s’écrit
`=` + `=`
