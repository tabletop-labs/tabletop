import Head from 'next/head'
import { Inter } from '@next/font/google'

const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  return (
    <>
      <Head>
        <title>Tabletop - Home</title>
      </Head>
      <main className={inter.className}>
        TODO... Home. All systems operational
      </main>
    </>
  )
}
