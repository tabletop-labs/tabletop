import "../styles/globals.css"
import Head from 'next/head'
import type { AppProps } from "next/app"
import Link from "next/link";

export default function App({ Component, pageProps }: AppProps) {
  return (
    <>
      <Head>
        <title>Tabletop - The modern data toolkit for infinite real-time streams</title>
        <meta name="description" content="The modern data toolkit for infinite real-time streams" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <MainNav />
      <Component {...pageProps} />
    </>
  )
}

function MainNav() {
  const q = ""
  const submit = () => { console.log("submit") }

  return (
    <>
      <div className="flex items-center justify-between border-b-2 pt-2 pb-2 h-20 bg-base-200">
        <div className="flex items-center ml-8">
          <div>
            <Link href="/">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-10 h-10">
                <path strokeLinecap="round" strokeLinejoin="round" d="M6.429 9.75L2.25 12l4.179 2.25m0-4.5l5.571 3 5.571-3m-11.142 0L2.25 7.5 12 2.25l9.75 5.25-4.179 2.25m0 0L21.75 12l-4.179 2.25m0 0l4.179 2.25L12 21.75 2.25 16.5l4.179-2.25m11.142 0l-5.571 3-5.571-3" />
              </svg>
            </Link>
          </div>
          <div>
            <form id="search-form" role="search" className="ml-4">
              <div className="form-control">
                <div className="input-group">
                  <input
                    placeholder="Search…"
                    id="q"
                    className="input input-bordered w-96"
                    aria-label="Search contacts"
                    type="search"
                    name="q"
                    defaultValue={q}
                    onChange={submit}
                  />
                  <button className="btn btn-square">
                    <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" /></svg>
                  </button>
                </div>
              </div>
            </form>
          </div>
        </div>

        <div className="mr-8">
          <nav className="grid">
            <ul>
              <li className="inline-block">
                <Link href="/topology">Topology</Link>
              </li>
              <li className="inline-block">
                &nbsp;|&nbsp;
                <Link href="/plugins">Plugins</Link>
              </li>
              <li className="inline-block">
                &nbsp;|&nbsp;
                <a href="http://redpanda.localhost">Redpanda</a>
              </li>
              <li className="inline-block">
                &nbsp;|&nbsp;
                <a href="http://grafana.localhost">Grafana</a>
              </li>
              <li className="inline-block">
                &nbsp;|&nbsp;
                <a href="http://prometheus.localhost">Prometheus</a>
              </li>
            </ul>
          </nav>
        </div>
      </div>
    </>
  );
}
