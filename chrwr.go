package chromedp_wrapper

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/i-b8o/chromedp_wrapper/scripts"
)

type Chrome struct {
	TimeOut int
}

func Init() (context.Context, context.CancelFunc) {
	opts := []chromedp.ExecAllocatorOption{chromedp.Flag("no-sandbox", true)}
	allocContext, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	return chromedp.NewContext(allocContext)
}

func InitHeadLess() (context.Context, context.CancelFunc) {
	opts := []chromedp.ExecAllocatorOption{chromedp.Flag("no-sandbox", true), chromedp.Flag("headless", false)}
	allocContext, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	return chromedp.NewContext(allocContext)
}

func NewChromeWrapper() *Chrome {
	return &Chrome{TimeOut: 60}
}

func (c *Chrome) SetTimeout(timeOut int) {
	c.TimeOut = timeOut
}

func openURL(url string, message *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.EvaluateAsDevTools(scripts.OpenURL(url), message),
	}
}

func (c *Chrome) OpenURL(ctxt context.Context, url string) error {
	var message string
	err := chromedp.Run(ctxt, RunWithTimeOut(&ctxt, 60, openURL(url, &message)))
	return err
}

func RunWithTimeOut(ctx *context.Context, timeout time.Duration, tasks chromedp.Tasks) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		timeoutContext, cancel := context.WithTimeout(ctx, timeout*time.Second)
		defer cancel()
		return tasks.Do(timeoutContext)
	}
}

func waitVisible(selector string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitVisible(selector, chromedp.ByQuery),
	}
}

func (c *Chrome) WaitVisible(ctxt context.Context, selector string) error {
	err := chromedp.Run(ctxt, RunWithTimeOut(&ctxt, 60, waitVisible(selector)))
	return err
}

func waitReady(selector string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitReady(selector, chromedp.ByQuery),
	}
}

func (c *Chrome) WaitReady(ctxt context.Context, selector string) error {
	err := chromedp.Run(ctxt, RunWithTimeOut(&ctxt, 60, waitReady(selector)))
	return err
}

func getString(jsString string, resultString *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.EvaluateAsDevTools(scripts.GetValue(jsString), resultString),
	}
}

func (c *Chrome) GetString(ctxt context.Context, jsString string) (string, error) {
	var resultString string
	err := chromedp.Run(ctxt, RunWithTimeOut(&ctxt, 60, getString(jsString, &resultString)))
	return resultString, err
}

func getStringsSlice(jsString string, resultSlice *[]string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.EvaluateAsDevTools(scripts.GetValue(jsString), resultSlice),
	}
}

func (c *Chrome) GetStringsSlice(ctxt context.Context, jsString string) ([]string, error) {
	var stringSlice []string
	err := chromedp.Run(ctxt, RunWithTimeOut(&ctxt, 60, getStringsSlice(jsString, &stringSlice)))
	return stringSlice, err
}

func getBool(jsBool string, resultBool *bool) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.EvaluateAsDevTools(scripts.GetValue(jsBool), resultBool),
	}
}

func (c *Chrome) GetBool(ctxt context.Context, jsBool string) (bool, error) {
	var resultBool bool
	err := chromedp.Run(ctxt, RunWithTimeOut(&ctxt, 60, getBool(jsBool, &resultBool)))
	return resultBool, err
}

func click(selector string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Sleep(1 * time.Second),
		chromedp.Click(selector, chromedp.ByQuery),
	}
}

func (c *Chrome) Click(ctxt context.Context, selector string) error {
	err := chromedp.Run(ctxt, RunWithTimeOut(&ctxt, 60, waitVisible(selector)))
	if err != nil {
		return err
	}
	return chromedp.Run(ctxt, RunWithTimeOut(&ctxt, 60, click(selector)))

}

func (c *Chrome) WaitLoaded(ctxt context.Context) error {
	var loaded bool
	loaded, err := c.GetBool(ctxt, `document.readyState !== 'ready' && document.readyState !== 'complete'`)
	if err != nil {
		return err
	}
	n := 0
	for loaded {
		if n > c.TimeOut {
			return fmt.Errorf("time is over: %d sec", c.TimeOut)
		}
		time.Sleep(1 * time.Second)
		loaded, err = c.GetBool(ctxt, `document.readyState !== 'ready' && document.readyState !== 'complete'`)
		if err != nil {
			return err
		}
		n++
	}
	return nil
}
