import React, { useState, useEffect } from 'react';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import axios from '../../axios';
import Course from "./Course"
import CourseForm from "./CourseForm"
import requests from '../../requests'

function CourseList({ title, fetchUrl }) {
    const [courses, setCourses] = useState([])
    const [showForm, setShowForm] = useState(false)

    async function fetchData() {
        const res = await axios.get(fetchUrl, {mode:'cors'});
        setCourses(res.data.data)
    }
    useEffect(() => {
        fetchData();
    }, [fetchUrl])

    function createNewCourse() {
        setShowForm(true)
    }
    function closeForm() {
        setShowForm(false)
    }
    const onSubmit = data => {
        axios.post(requests.courses, data, {mode:'cors'}).then(res => {
            toast.success("Tạo mới thành công! CourseId:" + res.data.data);
            closeForm();
            fetchData();
        })
            .catch(err => {
                if (err.response) {
                    toast.error(err.response.statusText + ": " + err.response.data.log)
                }
            })
    }
    return (
        <div>
            <ToastContainer
                position="top-right"
                autoClose={5000}
                hideProgressBar={false}
                newestOnTop={false}
                closeOnClick
                rtl={false}
                pauseOnFocusLoss
                draggable
                pauseOnHover
                theme="colored"
            />
            <div className="container mx-auto my-8">
                <div className="flex justify-between items-center">
                    <div className="text-2xl font-medium text-blue-900">{title}</div>
                    <button className="py-2 px-4 font-semibold rounded-lg shadow-md text-white bg-green-500 hover:bg-green-700"
                        onClick={createNewCourse}>
                        Tạo khoá học mới
                    </button>
                </div>

                <div>
                    {courses.map(course => (
                        <Course course={course} key={course.code}></Course>
                    ))}
                </div>
            </div>

            {showForm && <CourseForm closeForm={closeForm} onSubmit={onSubmit} />}

        </div>

    )
}

export default CourseList