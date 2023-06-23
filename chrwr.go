package chromedp_wrapper

import (
	"bufio"
	"context"

	"log"
	"os"

	"time"

	"github.com/chromedp/chromedp"
)

const (
	DefaultTimeout = 60 * time.Second
)

type Chrome struct {
	Timeout time.Duration
	Logger  *log.Logger
}

func NewChromeWrapper() *Chrome {
	return &Chrome{
		Timeout: DefaultTimeout,
		Logger:  log.New(os.Stderr, "", log.LstdFlags),
	}
}

func (c *Chrome) SetTimeout(timeout time.Duration) {
	c.Timeout = timeout
}

func (c *Chrome) SetLogger(logger *log.Logger) {
	c.Logger = logger
}

func (c *Chrome) OpenURL(ctx context.Context, url string) error {
	if _, err := os.Stat(url); err == nil {
		file, err := os.Open(url)
		if err != nil {
			return err
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if err := c.openURL(ctx, scanner.Text()); err != nil {
				return err
			}
		}
		return scanner.Err()
	}
	return c.openURL(ctx, url)
}

func (c *Chrome) openURL(ctx context.Context, url string) error {
	ctx, cancel := context.WithTimeout(ctx, c.Timeout)
	defer cancel()
	c.Logger.Printf("Opening URL: %s", url)
	err := chromedp.Run(ctx, RunWithTimeout(openURL(url), c.Timeout))
	if err != nil {
		c.Logger.Printf("Error opening URL %s: %v", url, err)
		return err
	}
	return nil
}

func (c *Chrome) WaitVisible(ctx context.Context, selector string) error {
	c.Logger.Printf("Waiting for element to be visible: %s", selector)
	ctx, cancel := context.WithTimeout(ctx, c.Timeout)
	defer cancel()
	err := chromedp.Run(ctx, RunWithTimeout(waitVisible(selector), c.Timeout))
	if err != nil {
		c.Logger.Printf("Element never became visible: %s", selector)
		return err
	}
	return nil
}

func (c *Chrome) WaitReady(ctx context.Context, selector string) error {
	c.Logger.Printf("Waiting for element to be ready: %s", selector)
	ctx, cancel := context.WithTimeout(ctx, c.Timeout)
	defer cancel()
	err := chromedp.Run(ctx, RunWithTimeout(waitReady(selector), c.Timeout))
	if err != nil {
		c.Logger.Printf("Element never became ready: %s", selector)
		return err
	}
	return nil
}

func (c *Chrome) GetString(ctx context.Context, jsString string) (string, error) {
	var resultString string
	c.Logger.Printf("Getting string value: %s", jsString)
	ctx, cancel := context.WithTimeout(ctx, c.Timeout)
	defer cancel()
	err := chromedp.Run(ctx, RunWithTimeout(getString(jsString, &resultString), c.Timeout))
	if err != nil {
		c.Logger.Printf("Error getting string value: %v", err)
		return "", err
	}
	return resultString, nil
}

func (c *Chrome) GetStringSlice(ctx context.Context, jsString string) ([]string, error) {
	var stringSlice []string
	c.Logger.Printf("Getting string slice value: %s", jsString)
	ctx, cancel := context.WithTimeout(ctx, c.Timeout)
	defer cancel()
	err := chromedp.Run(ctx, RunWithTimeout(getStringSlice(jsString, &stringSlice), c.Timeout))
	if err != nil {
		c.Logger.Printf("Error getting string slice value: %v", err)
		return nil, err
	}
	return stringSlice, nil
}

func (c *Chrome) GetBool(ctx context.Context, jsBool string) (bool, error) {
	var resultBool bool
	c.Logger.Printf("Getting boolean value: %s", jsBool)
	ctx, cancel := context.WithTimeout(ctx, c.Timeout)
	defer cancel()
	err := chromedp.Run(ctx, RunWithTimeout(getBool(jsBool, &resultBool), c.Timeout))
	if err != nil {
		c.Logger.Printf("Error getting boolean value: %v", err)
		return false, err
	}
	return resultBool, nil
}

func (c *Chrome) Click(ctx context.Context, selector string) error {
	if err := c.WaitVisible(ctx, selector); err != nil {
		return err
	}
	c.Logger.Printf("Clicking element: %s", selector)
	ctx, cancel := context.WithTimeout(ctx, c.Timeout)
	defer cancel()
	err := chromedp.Run(ctx, RunWithTimeout(click(selector), c.Timeout))
	if err != nil {
		c.Logger.Printf("Error clicking element: %v", err)
		return err
	}
	return nil
}

func RunWithTimeout(tasks chromedp.Tasks, timeout time.Duration) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return tasks.Do(ctx)
	}
}

func openURL(url string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
	}
}

func waitVisible(selector string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitVisible(selector),
	}
}

func waitReady(selector string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitReady(selector),
	}
}

func getString(jsString string, result *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Evaluate(jsString, result),
	}
}

func getStringSlice(jsString string, result *[]string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Evaluate(jsString, result),
	}
}

func getBool(jsBool string, result *bool) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Evaluate(jsBool, result),
	}
}

func click(selector string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Click(selector),
	}
}
