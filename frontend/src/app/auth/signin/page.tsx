import React from 'react'
import Register from '../components/login'
import Login from '../components/login'

type Props = {}

function Page({}: Props) {
  return (
    <div className='h-screen flex items-center justify-center bg-neutral-200'>
        <Login />
    </div>
  )
}

export default Page