import { useQuery, gql } from "@apollo/client"

const GET_USERID = gql`
  query{
    user{
      userID
      results{
        resultID
      }
    }
  }
`

function App() {
  const {loading,error,data} = useQuery(GET_USERID)

  if (loading) return <p>Loading...</p>;
  if  (error) return <p>Error : {error.message}</p>;
  return (
    <div className="h-80 w-80 bg-red-500 text-8xl"> 
      {data.user.userID}
    </div>
  )
}

export default App
