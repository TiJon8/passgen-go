
# Go PassGen (password/seed/key generator)

Модуль для генерации последовательностей из набора символов, можно использовать, как генератор сидов, ключей или паролей

Модуль предоставляет генерацию через пакет go *crypto/rand*, который определяет умеренно оптимальную надежность для генерации паролей и **настоятельно** рекомендуется использовать именно NewSourceCrypto API для паролей


## Usage

```bash
go get https://github.com/TiJon8/passgen-go
```

```go
import (
	"fmt"
	passgen "github.com/TiJon8/passgen-go/pkg/lib"
)

func main()  {
    // based on crypto
	sourceCrypto := passgen.NewSourceCrypto(18, "")
	password := sourceCrypto.Generate()
	fmt.Println(password)

    // based on pseudo
    source := "0123456789"
	sourcePseudo := passgen.NewSourcePseudo(83729104, source)
	randomGen := sourcePseudo.Generate(18)
	fmt.Println(randomGen)
}
```


## API Reference

#### NewSourcePseudo

Принимает seed int64 и строковая последовательность, из которых будет сгенерированно значение. Возвращает _randSourcePseudo_ на котором можно вызвать Generate(length), где length длина генерации

seed - это число которое выступает базисом для псевдоуслучайной генерации, **при одном и том же значении, геренация будет одна и та же!** 

по умолчанию seed = time.Now().UnixNano()

#### NewSourceCrypto

Принимает length uint64 - длина генерации и строковую последовательность, из которой будет сгенерированно значение. Возвращает _randSourceCrypto_ на котром можно вызвать Generate()

NewSourceCrypto использует crypto/rand, который предоставляет экземпляр криптографически защищенного генератора, на разных системах, api crypto/rand испольузет разные механизмы

Прежде чем использовать для генерации паролей, убедитесь, наксолько вы доверяете механизмам go

[Go Reader Type](https://pkg.go.dev/crypto/rand@go1.27.0#pkg-variables)


## Использовать как CLI

Можно использовать как вызовы cli

```bash
  go install github.com/TiJon8/passgen-go/cmd/passgen
```

#### Использование passgen

```bash
  passgen -mode=0 \
  -seed=1972241498 \
  -source="abcdefghijklmnopqrstuvwxyz" \
  -len=10 \
  -count=2
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `mode` | `int` | тип генерации, 0 - pseudo, 1 - crypto; default: 0 |
| `seed` | `int64` | cид для генерации; default: time.Now().UnixNano() |
| `source` | `string` | последовательность символов для результата; default: "QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm01234567890!" |
| `len` | `uint64` | **Required**. длина последовательности |
| `count` | `int64` | количество итераций; default: 1 |


## Feedback

Если вы наткнулись на данный репозиторий и у вас есть вдохновение и мотивация его улучшить, прошу, эксперементируйте :)

Буду рад предложениям и идеям в [telegram](https://t.me/stesnyashka)

