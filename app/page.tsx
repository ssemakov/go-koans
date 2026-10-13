import Link from "next/link";

const firstKoan = `numbers := []int{1, 2, 3}
alias := numbers
alias[0] = 9

fmt.Println(numbers[0])`;

export default function Home() {
  return (
    <>
      <a className="skip-link" href="#main">Skip to content</a>
      <div className="site-shell">
        <header className="site-header">
          <Link className="brand" href="/" aria-label="Go Koans home">
            <span className="brand-mark" aria-hidden="true">g.</span>
            go koans
          </Link>
          <a className="text-link" href="https://github.com/ssemakov/go-koans">
            View source <span aria-hidden="true">↗</span>
          </a>
        </header>

        <main id="main">
          <section className="hero" aria-labelledby="hero-title">
            <p className="eyebrow"><span aria-hidden="true" /> A practice in curiosity</p>
            <h1 id="hero-title">Small questions.<br /><em>Lasting insights.</em></h1>
            <p className="hero-description">
              Learn Go one koan at a time. Read a little code, question what
              happens, and discover the idea hiding underneath.
            </p>
            <a className="button" href="#first-koan">
              Try your first koan <span aria-hidden="true">↓</span>
            </a>
            <p className="hero-note">No setup. Just a moment of thought.</p>
          </section>

          <section className="koan" id="first-koan" aria-labelledby="koan-title">
            <div className="koan-introduction">
              <p className="eyebrow">Koan 001 / Slices</p>
              <h2 id="koan-title">When is a copy<br />still connected?</h2>
              <p>Two variables. One change. What does this Go program print?</p>
            </div>
            <div className="code-card">
              <div className="code-header">
                <span><span className="code-dot" aria-hidden="true" /> slices.go</span>
                <span>Go</span>
              </div>
              <pre><code>{firstKoan}</code></pre>
              <details className="answer">
                <summary>Reveal the answer</summary>
                <div className="answer-content">
                  <p className="answer-result">It prints <code>9</code>.</p>
                  <p>
                    Assigning a slice copies its header. Both slices still share
                    the same underlying array, so changing an element through
                    <code> alias</code> also changes what <code>numbers</code> sees.
                  </p>
                  <p>To get independent elements, make a new slice and use <code>copy</code>.</p>
                </div>
              </details>
            </div>
          </section>

          <section className="practice" aria-labelledby="practice-title">
            <p className="eyebrow">The approach</p>
            <h2 id="practice-title">Understanding starts with a pause.</h2>
            <div className="practice-grid">
              <article><span className="step">01</span><h3>Read closely</h3><p>Small examples make each detail worth noticing.</p></article>
              <article><span className="step">02</span><h3>Make a prediction</h3><p>Give the question a little thought before revealing the answer.</p></article>
              <article><span className="step">03</span><h3>Follow the insight</h3><p>Understand the reason, then try the idea in your own code.</p></article>
            </div>
          </section>
        </main>

        <footer className="site-footer">
          <span>go koans</span>
          <span>A little code. A little clarity.</span>
        </footer>
      </div>
    </>
  );
}
