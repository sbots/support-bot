import React from "react";
import { useForm } from "react-hook-form";
import { Input } from "../components/Input/Input";
import Link from "next/link";
// import { useMutation } from 'react-query'
// import Router from "next/router";
// import Image from 'next/image'

type Inputs = {
  usernameOrEmail: string;
  password: string;
};

const SignIn = () => {
  const { register, handleSubmit } = useForm<Inputs>();

  const onSubmit = (data: Inputs) => {
    console.log(data);
  };
  return (
    <div>
      <div className="min-h-screen bg-white flex">
        <div className="flex-1 flex flex-col justify-center py-12 px-4 sm:px-6 lg:flex-none lg:px-20 xl:px-24">
          <div className="mx-auto w-full max-w-sm lg:w-96">
            <div>
              <img
                className="h-12 w-auto"
                src="https://tailwindui.com/img/logos/workflow-mark-indigo-600.svg"
                alt="Workflow"
              />
              <h2 className="mt-6 text-3xl font-extrabold text-gray-900">
                Sign in to your account
              </h2>
            </div>

            <div className="mt-8">
              <div className="mt-6">
                <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
                  <div>
                    <Input
                      label="Email address"
                      type="email"
                      {...register("usernameOrEmail", {
                        required: true,
                      })}
                    />
                  </div>

                  <div className="space-y-1">
                    <Input
                      label="Password"
                      type="password"
                      className="space-y-1"
                      {...register("password", {
                        required: true,
                      })}
                    />
                  </div>

                  <div className="flex items-center justify-between">
                    <div className="text-sm">
                      <Link href="/forgot-password">
                        <a className="font-medium text-indigo-600 hover:text-indigo-500">
                          Forgot your password?
                        </a>
                      </Link>
                    </div>
                    <div className="text-sm">
                      <Link href="/sign-up">
                        <a className="font-medium text-indigo-600 hover:text-indigo-500">
                          Dont have an account?
                        </a>
                      </Link>
                    </div>
                  </div>

                  <div>
                    <button
                      type="submit"
                      className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                    >
                      Sign in
                    </button>
                  </div>
                </form>
              </div>
            </div>
          </div>
        </div>
        <div className="hidden lg:block relative w-0 flex-1">
          <img
            className="absolute inset-0 h-full w-full object-cover"
            src="https://images.unsplash.com/photo-1505904267569-f02eaeb45a4c?ixlib=rb-1.2.1&amp;ixid=eyJhcHBfaWQiOjEyMDd9&amp;auto=format&amp;fit=crop&amp;w=1908&amp;q=80"
            alt=""
          />
        </div>
      </div>
    </div>
  );
};
export default SignIn;
