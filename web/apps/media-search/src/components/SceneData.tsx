// Copyright 2024 Google, LLC
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     https://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import {Grid2, Typography} from "@mui/material";
import {Scene} from "../shared/model";

const SceneData = ({scene}: { scene: Scene }) => {

    const formatScript = (val: string): string => {
        return val.replace("\n", "<br/>")
    }

    return (
        <>
            <Grid2 size={2}>{scene.sequence}</Grid2>
            <Grid2 size={5}>{scene.start}</Grid2>
            <Grid2 size={5}>{scene.end}</Grid2>
            <Grid2 size={12} sx={{textAlign: 'left'}}><Typography component="div" variant="body2">
                <div dangerouslySetInnerHTML={{__html: formatScript(scene.script)}}/>
            </Typography></Grid2>
        </>)
};

export default SceneData