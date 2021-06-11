import './App.css';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import {useState} from "react";

// I don't pretend this is idiomatic React 😉

const packsApiUrl = 'https://packs-example.projects.robertcurran.uk';


const App = () => {
    const [requestedCount, setRequestedCount] = useState(25);
    const [packingList, setPackingList] = useState();

    const getPackingListCalculationResult = async () => {

        const response = await fetch(`${packsApiUrl}/packs?requestedCount=${requestedCount}`, {
            method: 'GET',
            mode: "cors",
            headers: {
                "Content-Type": "application/json",
            },
        });
        const packingList = await response.json();
        console.log('Successfully acquired packingList:');
        console.log(JSON.stringify(packingList));
        setPackingList(packingList);
    }

    return (
        <div className="App">
            <header className="App-header">
                Packing List
            </header>
            <article>
                <p>This utility calculates the packing list required for a given number of widgets</p>


                <div style={packUtilityStyle}>
                    <div>
                        <TextField id="requestedCount" label="Widget Count"
                                   onChange={event => setRequestedCount(event.target.value)} value={requestedCount}/>

                        <Button onClick={getPackingListCalculationResult} variant="contained" color="primary">
                            Calculate
                        </Button>
                    </div>

                    <div style={resultContainer}>
                        Result:
                        <pre>
                    {packingList &&
                    JSON.stringify(packingList)
                    }
                        </pre>
                    </div>

                </div>

            </article>
        </div>
    );
}

const packUtilityStyle = {
    display: 'flex',
    justifyContent: 'space-evenly',
    flexDirection: 'column',
    margin: 'auto',
    maxWidth: '40vw',
    marginTop: '10vh',
    padding: '25px'
}

const resultContainer = {
    display: 'flex',
    padding: '25px'
}

export default App;
