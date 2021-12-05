import CourseList from "./components/Course/CourseList.jsx"
import requests from './requests'

function App() {
  return (
    <div>
      <div className="text-2xl font-bold text-green-500 text-center bg-white py-4 cursor-pointer">HCMUSMODS</div>
      <CourseList
        title="Khoá học hiện có"
        fetchUrl={requests.courses}
      />
    </div>
  );
}

export default App;
