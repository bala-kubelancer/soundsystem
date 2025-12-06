package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 👉 Only change this for your ArgoCD demo
	soundValue := "DOLBY ATMOS" // e.g. "dolby", "atmos", "7.1", etc.

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		fmt.Fprintf(w, `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <title>GitOps Demo App</title>
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <style>
    :root {
      --bg: #050816;
      --accent: #ff6bcb;
      --accent-soft: #7b5cff;
      --text: #f9fafb;
      --muted: #9ca3af;
    }

    * {
      box-sizing: border-box;
      margin: 0;
      padding: 0;
    }

    body {
      min-height: 100vh;
      background: radial-gradient(circle at top, #1f2937 0, #050816 45%%, #02010a 100%%);
      color: var(--text);
      font-family: system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: space-between;
      padding: 24px;
    }

    header, footer {
      width: 100%%;
      max-width: 720px;
      text-align: center;
    }

    header h1 {
      font-size: 26px;
      letter-spacing: 0.18em;
      text-transform: uppercase;
      font-weight: 600;
      padding-bottom: 12px;
      border-bottom: 1px solid rgba(148, 163, 184, 0.4);
    }

    footer p {
      font-size: 13px;
      color: var(--muted);
      padding-top: 10px;
      border-top: 1px solid rgba(148, 163, 184, 0.25);
    }

    main {
      flex: 1;
      width: 100%%;
      display: flex;
      align-items: center;
      justify-content: center;
    }

    .panel {
      width: 100%%;
      max-width: 520px;
      padding: 24px 26px 26px;
      background: rgba(15, 23, 42, 0.9);
      border-radius: 20px;
      border: 1px solid rgba(148, 163, 184, 0.5);
      box-shadow:
        0 18px 45px rgba(0, 0, 0, 0.65),
        0 0 35px rgba(123, 92, 255, 0.2);
      backdrop-filter: blur(12px);
    }

    .boxes {
      display: flex;
      flex-direction: column;
      gap: 12px;
      margin-bottom: 18px;
    }

    .box {
      padding: 16px 18px;
      border-radius: 14px;
      border: 1px solid rgba(148, 163, 184, 0.6);
      background: radial-gradient(circle at top left, rgba(255, 107, 203, 0.12), rgba(15, 23, 42, 0.98));
      text-align: center;
    }

    .label-box {
      font-size: 18px;
      font-weight: 500;
      color: var(--muted);
      text-transform: uppercase;
      letter-spacing: 0.16em;
    }

    .value-box {
      font-size: 30px;
      font-weight: 800;
      letter-spacing: 0.2em;
      text-transform: uppercase;
      color: var(--accent);
      background: radial-gradient(circle at top, rgba(123, 92, 255, 0.22), rgba(15, 23, 42, 1));
      box-shadow: 0 0 16px rgba(255, 107, 203, 0.35);
      animation: valuePulse 2.4s ease-in-out infinite;
    }

    /* small equalizer */
    .equalizer {
      margin-top: 10px;
      display: flex;
      justify-content: center;
      align-items: flex-end;
      gap: 5px;
      height: 40px;
    }

    .bar {
      width: 4px;
      border-radius: 999px;
      background: linear-gradient(to top, rgba(255, 107, 203, 0.1), var(--accent-soft));
      box-shadow: 0 0 8px rgba(123, 92, 255, 0.7);
      transform-origin: bottom;
      animation: barMove 1s ease-in-out infinite;
    }

    .bar:nth-child(2) { animation-duration: 0.7s; }
    .bar:nth-child(3) { animation-duration: 1.1s; }
    .bar:nth-child(4) { animation-duration: 0.8s; }
    .bar:nth-child(5) { animation-duration: 1.2s; }

    @keyframes valuePulse {
      0%%, 100%% { transform: scale(1); box-shadow: 0 0 10px rgba(255, 107, 203, 0.25); }
      50%%        { transform: scale(1.03); box-shadow: 0 0 20px rgba(255, 107, 203, 0.5); }
    }

    @keyframes barMove {
      0%%   { transform: scaleY(0.3); }
      25%%  { transform: scaleY(0.9); }
      50%%  { transform: scaleY(0.45); }
      75%%  { transform: scaleY(1.0); }
      100%% { transform: scaleY(0.4); }
    }

    @media (max-width: 600px) {
      header h1 {
        font-size: 18px;
        letter-spacing: 0.14em;
      }
      .panel {
        padding: 18px 16px 20px;
      }
      .label-box {
        font-size: 15px;
      }
      .value-box {
        font-size: 22px;
        letter-spacing: 0.16em;
      }
    }
  </style>
</head>
<body>
  <header>
    <h1>GitOps Demo App</h1>
  </header>

  <main>
    <div class="panel">
      <div class="boxes">
        <div class="box label-box">Sound System</div>
        <div class="box value-box">%s</div>
      </div>

      <div class="equalizer">
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
      </div>
    </div>
  </main>

  <footer>
    <p>Powered by ArgoCD & GitOps · Live DJ View</p>
  </footer>
</body>
</html>
`, soundValue)
	})

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
