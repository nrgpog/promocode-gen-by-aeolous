package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/chromedp/chromedp"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)

func generateRandomName() string {
	names := []string{"Alex", "John", "Sarah", "Mike", "Emma", "David", "Lisa", "Chris", "Anna", "Tom"}
	return names[rand.Intn(len(names))]
}

func generateRandomDisplayName() string {
	chars := "abcdefghijklmnopqrstuvwxyz0123456789"
	name := make([]byte, 10)
	for i := range name {
		name[i] = chars[rand.Intn(len(chars))]
	}
	return string(name)
}

func generateRandomPassword() string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	password := make([]byte, 12)
	for i := range password {
		password[i] = chars[rand.Intn(len(chars))]
	}
	return string(password)
}

func main() {
	log.SetOutput(io.Discard)
	
	rand.Seed(time.Now().UnixNano())
	
	fmt.Printf("%s🚀 Iniciando proceso de automatización...%s\n", ColorCyan, ColorReset)
	
	tempDir, _ := os.MkdirTemp("", "chrome-temp")
	defer os.RemoveAll(tempDir)

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.Flag("exclude-switches", "enable-automation"),
		chromedp.Flag("disable-extensions-except", ""),
		chromedp.Flag("disable-extensions", false),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
		chromedp.WindowSize(1920, 1080),
		chromedp.UserDataDir(tempDir),
		chromedp.Flag("disable-web-security", false),
		chromedp.Flag("disable-features", "VizDisplayCompositor,TranslateUI"),
		chromedp.Flag("disable-ipc-flooding-protection", true),
	)

	allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer allocCancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	ctx, timeoutCancel := context.WithTimeout(ctx, 300*time.Second)
	defer timeoutCancel()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Printf("%s\n🛑 Proceso interrumpido, cerrando...%s\n", ColorRed, ColorReset)
		cancel()
		allocCancel()
		os.Exit(0)
	}()

	fmt.Printf("%s🌐 Navegando a Epic Games...%s\n", ColorBlue, ColorReset)

	err := chromedp.Run(ctx,
		chromedp.Navigate("https://store.epicgames.com/en-US/p/discord--discord-nitro"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s✅ Página cargada correctamente%s\n", ColorGreen, ColorReset)
			return nil
		}),
		chromedp.Evaluate(`
			Object.defineProperty(navigator, 'webdriver', {
				get: () => undefined,
			});
			Object.defineProperty(navigator, 'plugins', {
				get: () => [1, 2, 3, 4, 5],
			});
			Object.defineProperty(navigator, 'languages', {
				get: () => ['en-US', 'en'],
			});
			window.chrome = {
				runtime: {},
			};
		`, nil),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s🔧 Configuración anti-detección aplicada%s\n", ColorPurple, ColorReset)
			return nil
		}),
		chromedp.Sleep(5*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s🎂 Iniciando verificación de edad...%s\n", ColorYellow, ColorReset)
			return nil
		}),
		chromedp.WaitVisible("#month_toggle"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s✅ Elemento mes encontrado%s\n", ColorGreen, ColorReset)
			return nil
		}),
		chromedp.Click("#month_toggle"),
		chromedp.Sleep(1*time.Second),
		chromedp.WaitVisible("#month_menu > li:nth-child(4)"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s✅ Seleccionando mes%s\n", ColorGreen, ColorReset)
			return nil
		}),
		chromedp.Click("#month_menu > li:nth-child(4)"),
		chromedp.WaitVisible("#day_toggle"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s✅ Elemento día encontrado%s\n", ColorGreen, ColorReset)
			return nil
		}),
		chromedp.Click("#day_toggle"),
		chromedp.Sleep(1*time.Second),
		chromedp.WaitVisible("#day_menu > li:nth-child(4)"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s✅ Seleccionando día%s\n", ColorGreen, ColorReset)
			return nil
		}),
		chromedp.Click("#day_menu > li:nth-child(4)"),
		chromedp.WaitVisible("#year_toggle"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s✅ Elemento año encontrado%s\n", ColorGreen, ColorReset)
			return nil
		}),
		chromedp.Click("#year_toggle"),
		chromedp.Sleep(1*time.Second),
		chromedp.WaitVisible("#year_menu > li:nth-child(41)"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s✅ Seleccionando año%s\n", ColorGreen, ColorReset)
			return nil
		}),
		chromedp.Click("#year_menu > li:nth-child(41)"),
		chromedp.WaitVisible("#btn_age_continue"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s✅ Botón continuar edad encontrado%s\n", ColorGreen, ColorReset)
			return nil
		}),
		chromedp.Click("#btn_age_continue"),
		chromedp.Sleep(2*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s🛒 Buscando botón de compra...%s\n", ColorYellow, ColorReset)
			return nil
		}),
		chromedp.WaitVisible(`button[data-testid="purchase-cta-button"]`),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s✅ Botón de compra encontrado%s\n", ColorGreen, ColorReset)
			return nil
		}),
		chromedp.Click(`button[data-testid="purchase-cta-button"]`),
		chromedp.Sleep(2*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s📝 Buscando opción de registro...%s\n", ColorYellow, ColorReset)
			return nil
		}),
		chromedp.WaitVisible("#to-register"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s✅ Opción de registro encontrada%s\n", ColorGreen, ColorReset)
			return nil
		}),
		chromedp.Click("#to-register"),
		chromedp.Sleep(2*time.Second),
		chromedp.WaitVisible("#no"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s✅ Seleccionando 'No' para cuenta existente%s\n", ColorGreen, ColorReset)
			return nil
		}),
		chromedp.Click("#no"),
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s🎂 Configurando fecha de nacimiento...%s\n", ColorYellow, ColorReset)
			return nil
		}),
		chromedp.WaitVisible("#month"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s✅ Campo mes encontrado%s\n", ColorGreen, ColorReset)
			return nil
		}),
		chromedp.Click("#month"),
		chromedp.Sleep(1*time.Second),
		chromedp.WaitVisible(`#\:rd\: > li:nth-child(6)`),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s✅ Seleccionando junio%s\n", ColorGreen, ColorReset)
			return nil
		}),
		chromedp.Click(`#\:rd\: > li:nth-child(6)`),
		chromedp.Sleep(1*time.Second),
		chromedp.WaitVisible("#day"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s✅ Campo día encontrado%s\n", ColorGreen, ColorReset)
			return nil
		}),
		chromedp.Click("#day"),
		chromedp.Sleep(1*time.Second),
		chromedp.WaitVisible(`#\:rf\: > li:nth-child(5)`),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s✅ Seleccionando día 5%s\n", ColorGreen, ColorReset)
			return nil
		}),
		chromedp.Click(`#\:rf\: > li:nth-child(5)`),
		chromedp.Sleep(1*time.Second),
		chromedp.WaitVisible("#year"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s✅ Campo año encontrado%s\n", ColorGreen, ColorReset)
			return nil
		}),
		chromedp.Click("#year"),
		chromedp.SendKeys("#year", "1999"),
		chromedp.Sleep(1*time.Second),
		chromedp.WaitVisible("#continue"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s✅ Botón continuar encontrado%s\n", ColorGreen, ColorReset)
			return nil
		}),
		chromedp.Click("#continue"),
		chromedp.Sleep(2*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s👤 Llenando información personal...%s\n", ColorYellow, ColorReset)
			return nil
		}),
		chromedp.WaitVisible("#name"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s✅ Campo nombre encontrado%s\n", ColorGreen, ColorReset)
			return nil
		}),
		chromedp.Click("#name"),
		chromedp.SendKeys("#name", generateRandomName()),
		chromedp.Sleep(1*time.Second),
		chromedp.WaitVisible("#lastName"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s✅ Campo apellido encontrado%s\n", ColorGreen, ColorReset)
			return nil
		}),
		chromedp.Click("#lastName"),
		chromedp.SendKeys("#lastName", generateRandomName()),
		chromedp.Sleep(1*time.Second),
		chromedp.WaitVisible("#displayName"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s✅ Campo nombre de usuario encontrado%s\n", ColorGreen, ColorReset)
			return nil
		}),
		chromedp.Click("#displayName"),
		chromedp.SendKeys("#displayName", generateRandomDisplayName()),
		chromedp.Sleep(1*time.Second),
		chromedp.WaitVisible("#password"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s✅ Campo contraseña encontrado%s\n", ColorGreen, ColorReset)
			return nil
		}),
		chromedp.Click("#password"),
		chromedp.SendKeys("#password", generateRandomPassword()),
		chromedp.Sleep(1*time.Second),
		chromedp.WaitVisible("#tos"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s✅ Checkbox términos encontrado%s\n", ColorGreen, ColorReset)
			return nil
		}),
		chromedp.Click("#tos"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s📧 Esperando a la introducción del correo...%s\n", ColorCyan, ColorReset)
			for {
				var buttonClass string
				var disabled string
				err1 := chromedp.AttributeValue("#btn-submit", "class", &buttonClass, nil).Do(ctx)
				err2 := chromedp.AttributeValue("#btn-submit", "disabled", &disabled, nil).Do(ctx)
				
				if err1 == nil {
					hasDisabledClass := strings.Contains(buttonClass, "Mui-disabled")
					hasDisabledAttr := err2 == nil && disabled != ""
					
					if !hasDisabledClass && !hasDisabledAttr {
						fmt.Printf("%s✅ Email detectado, continuando...%s\n", ColorGreen, ColorReset)
						break
					}
				}
				time.Sleep(2 * time.Second)
			}
			return nil
		}),
		chromedp.WaitVisible("#btn-submit"),
		chromedp.Click("#btn-submit"),
		chromedp.Sleep(2*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s📧 Esperando verificación de email...%s\n", ColorYellow, ColorReset)
			for {
				var buttonClass string
				var disabled string
				err1 := chromedp.AttributeValue("#continue", "class", &buttonClass, nil).Do(ctx)
				err2 := chromedp.AttributeValue("#continue", "disabled", &disabled, nil).Do(ctx)
				
				if err1 == nil {
					hasDisabledClass := strings.Contains(buttonClass, "Mui-disabled")
					hasDisabledAttr := err2 == nil && disabled != ""
					
					if !hasDisabledClass && !hasDisabledAttr {
						fmt.Printf("%s✅ Botón de verificación habilitado%s\n", ColorGreen, ColorReset)
						break
					}
				}
				time.Sleep(2 * time.Second)
			}
			return nil
		}),
		chromedp.WaitVisible("#continue"),
		chromedp.Click("#continue"),
		chromedp.Sleep(2*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s🔗 Buscando botón de vinculación...%s\n", ColorYellow, ColorReset)
			return nil
		}),
		chromedp.WaitVisible("#link-success"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s✅ Botón de vinculación encontrado%s\n", ColorGreen, ColorReset)
			return nil
		}),
		chromedp.Click("#link-success"),
		chromedp.Sleep(2*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s🛡️ Buscando checkbox de privacidad...%s\n", ColorYellow, ColorReset)
			return nil
		}),
		chromedp.WaitVisible("#payment-summaries__scroll-container > div.payment-developer-privacy > div > div"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s✅ Checkbox de privacidad encontrado%s\n", ColorGreen, ColorReset)
			return nil
		}),
		chromedp.Click("#payment-summaries__scroll-container > div.payment-developer-privacy > div > div"),
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s💳 Buscando botón de pedido...%s\n", ColorYellow, ColorReset)
			return nil
		}),
		chromedp.WaitVisible("#purchase-app > div > div > div > div.payment-summaries > div.payment-confirm-container > div"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s✅ Botón de pedido encontrado%s\n", ColorGreen, ColorReset)
			return nil
		}),
		chromedp.Click("#purchase-app > div > div > div > div.payment-summaries > div.payment-confirm-container > div"),
		chromedp.Sleep(2*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s⏳ Esperando confirmación de pedido...%s\n", ColorYellow, ColorReset)
			return nil
		}),
		chromedp.WaitVisible(`[data-testid="checkout-success-title"]`),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Printf("%s🎉 ¡Pedido completado exitosamente!%s\n", ColorGreen, ColorReset)
			return nil
		}),
		chromedp.Sleep(4*time.Second),
	)

	if err != nil {
		fmt.Printf("%s❌ Error: %v%s\n", ColorRed, err, ColorReset)
		log.Fatal(err)
	}

	fmt.Printf("%s✅ Proceso completado exitosamente%s\n", ColorGreen, ColorReset)
} 