// Copyright (c) 2017-2018 Alexander Eichhorn
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package main

import (
	"encoding/json"
	"testing"

	_ "github.com/mholt/go-xmp/models"
	"github.com/mholt/go-xmp/xmp"
)

// Benchmarks
func BenchmarkUnmarshalXMP_5kB(B *testing.B)  { runUnmarshalXmpBench(B, []byte(data_5kb)) }
func BenchmarkMarshalXMP_5kB(B *testing.B)    { runMarshalXmpBench(B, []byte(data_5kb)) }
func BenchmarkMarshalJSON_5kB(B *testing.B)   { runMarshalJsonBench(B, []byte(data_5kb)) }
func BenchmarkUnmarshalJSON_5kB(B *testing.B) { runUnmarshalJsonBench(B, []byte(data_5kb)) }

func BenchmarkUnmarshalXMP_85kB(B *testing.B)  { runUnmarshalXmpBench(B, []byte(data_85kb)) }
func BenchmarkMarshalXMP_85kB(B *testing.B)    { runMarshalXmpBench(B, []byte(data_85kb)) }
func BenchmarkMarshalJSON_85kB(B *testing.B)   { runMarshalJsonBench(B, []byte(data_85kb)) }
func BenchmarkUnmarshalJSON_85kB(B *testing.B) { runUnmarshalJsonBench(B, []byte(data_85kb)) }

func runUnmarshalXmpBench(B *testing.B, data []byte) {
	xmp.SetLogLevel(xmp.LogLevelError)
	for i := 0; i < B.N; i++ {
		d := &xmp.Document{}
		if err := xmp.Unmarshal(data, d); err != nil {
			B.Fatal(err)
		}
		d.Close()
	}
}

func runMarshalXmpBench(B *testing.B, data []byte) {
	xmp.SetLogLevel(xmp.LogLevelError)
	d := &xmp.Document{}
	if err := xmp.Unmarshal(data, d); err != nil {
		B.Fatal(err)
	}
	B.ResetTimer()
	var (
		err error
	)
	for i := 0; i < B.N; i++ {
		buf, err = xmp.Marshal(d)
		if err != nil {
			B.Fatal(err)
		}
	}
	d.Close()
}

func runMarshalJsonBench(B *testing.B, data []byte) {
	xmp.SetLogLevel(xmp.LogLevelError)
	d := &xmp.Document{}
	if err := xmp.Unmarshal(data, d); err != nil {
		B.Fatal(err)
	}
	B.ResetTimer()
	var (
		err error
	)
	for i := 0; i < B.N; i++ {
		buf, err = json.MarshalIndent(d, "", "  ")
		if err != nil {
			B.Fatal(err)
		}
	}
	d.Close()
}

func runUnmarshalJsonBench(B *testing.B, data []byte) {
	xmp.SetLogLevel(xmp.LogLevelError)
	d := &xmp.Document{}
	if err := xmp.Unmarshal(data, d); err != nil {
		B.Fatal(err)
	}
	buf, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		B.Fatal(err)
	}
	B.ResetTimer()
	for i := 0; i < B.N; i++ {
		dx := &xmp.Document{}
		if err = json.Unmarshal(buf, dx); err != nil {
			B.Fatal(err)
		}
	}
	d.Close()
}

const data_85kb = `<?xpacket begin="" id="W5M0MpCehiHzreSzNTczkc9d"?>
<x:xmpmeta xmlns:x="adobe:ns:meta/" x:xmptk="Adobe XMP Core 5.2-c003 61.141987, 2011/02/22-12:03:51        ">
 <rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#">
  <rdf:Description rdf:about=""
    xmlns:xmp="http://ns.adobe.com/xap/1.0/"
    xmlns:creatorAtom="http://ns.adobe.com/creatorAtom/1.0/"
    xmlns:xmpDM="http://ns.adobe.com/xmp/1.0/DynamicMedia/"
    xmlns:stDim="http://ns.adobe.com/xap/1.0/sType/Dimensions#"
    xmlns:xmpMM="http://ns.adobe.com/xap/1.0/mm/"
    xmlns:stRef="http://ns.adobe.com/xap/1.0/sType/ResourceRef#"
    xmlns:stEvt="http://ns.adobe.com/xap/1.0/sType/ResourceEvent#"
    xmlns:dc="http://purl.org/dc/elements/1.1/"
    xmlns:photoshop="http://ns.adobe.com/photoshop/1.0/"
    xmlns:tiff="http://ns.adobe.com/tiff/1.0/"
    xmlns:exif="http://ns.adobe.com/exif/1.0/"
   xmp:CreatorTool="Adobe Premiere Pro 5.5"
   xmp:CreateDate="2012-09-01T18:44:56.0048+02:00"
   xmp:MetadataDate="2012-09-01T18:53:49+02:00"
   xmp:ModifyDate="2012-09-01T18:53:49.00156+02:00"
   xmpDM:videoFrameRate="25.000000"
   xmpDM:videoFieldOrder="Upper"
   xmpDM:videoPixelAspectRatio="17280/17280"
   xmpDM:audioSampleRate="48000"
   xmpDM:audioSampleType="Compressed"
   xmpDM:audioChannelType="Stereo"
   xmpDM:startTimeScale="25"
   xmpDM:startTimeSampleSize="1"
   xmpMM:OriginalDocumentID="xmp.did:2E0FA19553F4E1118E5EF593D63F81DA"
   xmpMM:InstanceID="xmp.iid:330FA19553F4E1118E5EF593D63F81DA"
   xmpMM:DocumentID="xmp.did:330FA19553F4E1118E5EF593D63F81DA"
   dc:format="MPEG2">
   <creatorAtom:windowsAtom
    creatorAtom:extension=".prproj"
    creatorAtom:invocationFlags="/L"
    creatorAtom:uncProjectPath="\\?\E:\alyaspot\Posnetki\Nova mapa\1080i.prproj"/>
   <creatorAtom:macAtom
    creatorAtom:applicationCode="1347449455"
    creatorAtom:invocationAppleEvent="1129468018"/>
   <xmpDM:projectRef
    xmpDM:type="movie"/>
   <xmpDM:videoFrameSize
    stDim:w="1920"
    stDim:h="1080"
    stDim:unit="pixel"/>
   <xmpDM:startTimecode
    xmpDM:timeFormat="25Timecode"
    xmpDM:timeValue="00:00:00:00"/>
   <xmpDM:altTimecode
    xmpDM:timeValue="00:00:00:00"
    xmpDM:timeFormat="25Timecode"/>
   <xmpMM:Ingredients>
    <rdf:Bag>
     <rdf:li
      stRef:instanceID="xmp.iid:2C0FA19553F4E1118E5EF593D63F81DA"
      stRef:documentID="xmp.did:8254602F53F4E11188E1CBA262DF7ED4"
      stRef:fromPart="time:0d11705057280000f254016000000"
      stRef:toPart="time:0d11705057280000f254016000000"
      stRef:filePath="trailer_maskiran_za_premiere.mov"
      stRef:maskMarkers="None"/>
     <rdf:li
      stRef:instanceID="xmp.iid:661D67351FF4E111A01FA07C9D969101"
      stRef:documentID="xmp.did:641D67351FF4E111A01FA07C9D969101"
      stRef:fromPart="time:386104320000f254016000000d11420559360000f254016000000"
      stRef:toPart="time:0d11420559360000f254016000000"
      stRef:filePath="Alya - Moja pesem trailer_reverb.wav"
      stRef:maskMarkers="None"/>
    </rdf:Bag>
   </xmpMM:Ingredients>
   <xmpMM:History>
    <rdf:Seq>
     <rdf:li
      stEvt:action="created"
      stEvt:instanceID="xmp.iid:2E0FA19553F4E1118E5EF593D63F81DA"
      stEvt:when="2012-09-01T18:48:53+02:00"
      stEvt:softwareAgent="Adobe Premiere Pro 5.5"/>
     <rdf:li
      stEvt:action="saved"
      stEvt:instanceID="xmp.iid:330FA19553F4E1118E5EF593D63F81DA"
      stEvt:when="2012-09-01T18:53:49+02:00"
      stEvt:softwareAgent="Adobe Premiere Pro 5.5"/>
    </rdf:Seq>
   </xmpMM:History>
   <xmpMM:Pantry>
    <rdf:Bag>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.comp"
       xmpMM:InstanceID="xmp.iid:13F95C2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">Comp 2</rdf:li>
       </rdf:Alt>
      </dc:title>
      <xmpMM:Ingredients>
       <rdf:Bag>
        <rdf:li
         stRef:instanceID="xmp.iid:C6355C2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d76800f25600"
         stRef:toPart="time:0d76800f25600"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:C7355C2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d76800f25600"
         stRef:toPart="time:0d76800f25600"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:C8355C2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d76800f25600"
         stRef:toPart="time:0d76800f25600"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:C9355C2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d76800f25600"
         stRef:toPart="time:0d76800f25600"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:CA355C2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d76800f25600"
         stRef:toPart="time:0d76800f25600"
         stRef:maskMarkers="None"/>
       </rdf:Bag>
      </xmpMM:Ingredients>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:14F95C2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">Comp 2</rdf:li>
       </rdf:Alt>
      </dc:title>
      <xmpMM:Ingredients>
       <rdf:Bag>
        <rdf:li
         stRef:instanceID="xmp.iid:13F95C2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d76800f25600"
         stRef:toPart="time:0d76800f25600"
         stRef:maskMarkers="None"/>
       </rdf:Bag>
      </xmpMM:Ingredients>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:15F95C2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">Black Solid 3</rdf:li>
       </rdf:Alt>
      </dc:title>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:16F95C2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">QView Production</rdf:li>
       </rdf:Alt>
      </dc:title>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       xmp:CreateDate="2012-08-12T16:37:11Z"
       xmp:ModifyDate="2012-08-22T16:13:13Z"
       xmp:MetadataDate="2012-08-22T18:13:28+02:00"
       xmpDM:videoAlphaMode="none"
       xmpDM:audioSampleRate="48000"
       xmpDM:audioSampleType="Compressed"
       xmpDM:audioChannelType="Stereo"
       xmpDM:videoFrameRate="25.000000"
       xmpMM:InstanceID="xmp.iid:1705054974ECE111B9A38987502929D1"
       xmpMM:DocumentID="xmp.did:5855F98372ECE111B9A38987502929D1"
       xmpMM:OriginalDocumentID="xmp.did:5855F98372ECE111B9A38987502929D1">
      <xmpDM:videoFrameSize
       stDim:w="1920"
       stDim:h="1080"
       stDim:unit="pixel"/>
      <xmpDM:duration
       xmpDM:value="1468000"
       xmpDM:scale="1/25000"/>
      <xmpMM:History>
       <rdf:Seq>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:5955F98372ECE111B9A38987502929D1"
         stEvt:when="2012-08-22T18:13:13+02:00"
         stEvt:softwareAgent="Adobe Premiere Pro 5.5"
         stEvt:changed="/metadata;/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:1705054974ECE111B9A38987502929D1"
         stEvt:when="2012-08-22T18:13:28+02:00"
         stEvt:softwareAgent="Adobe Premiere Pro 5.5"
         stEvt:changed="/metadata"/>
       </rdf:Seq>
      </xmpMM:History>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:17F95C2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">presents</rdf:li>
       </rdf:Alt>
      </dc:title>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:18F95C2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">White Solid 2</rdf:li>
       </rdf:Alt>
      </dc:title>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:19F95C2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">bg2.jpg</rdf:li>
       </rdf:Alt>
      </dc:title>
      <xmpMM:Ingredients>
       <rdf:Bag>
        <rdf:li
         stRef:instanceID="xmp.iid:5933D21264F0E11180B8D0D5AB388124"
         stRef:documentID="xmp.did:5833D21264F0E11180B8D0D5AB388124"
         stRef:fromPart="time:0d72704f25600"
         stRef:toPart="time:0d72704f25600"
         stRef:maskMarkers="All"/>
       </rdf:Bag>
      </xmpMM:Ingredients>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.comp"
       xmpMM:InstanceID="xmp.iid:1AF95C2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">Comp 3</rdf:li>
       </rdf:Alt>
      </dc:title>
      <xmpMM:Ingredients>
       <rdf:Bag>
        <rdf:li
         stRef:instanceID="xmp.iid:15F95C2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d72704f25600"
         stRef:toPart="time:0d72704f25600"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:16F95C2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d72704f25600"
         stRef:toPart="time:0d72704f25600"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:17F95C2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d72704f25600"
         stRef:toPart="time:0d72704f25600"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:18F95C2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d72704f25600"
         stRef:toPart="time:0d72704f25600"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:19F95C2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d72704f25600"
         stRef:toPart="time:0d72704f25600"
         stRef:maskMarkers="None"/>
       </rdf:Bag>
      </xmpMM:Ingredients>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:1BF95C2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">Comp 3</rdf:li>
       </rdf:Alt>
      </dc:title>
      <xmpMM:Ingredients>
       <rdf:Bag>
        <rdf:li
         stRef:instanceID="xmp.iid:1AF95C2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d72704f25600"
         stRef:toPart="time:0d72704f25600"
         stRef:maskMarkers="None"/>
       </rdf:Bag>
      </xmpMM:Ingredients>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:1CF95C2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">Alya</rdf:li>
       </rdf:Alt>
      </dc:title>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:1DF95C2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">title.jpg</rdf:li>
       </rdf:Alt>
      </dc:title>
      <xmpMM:Ingredients>
       <rdf:Bag>
        <rdf:li
         stRef:instanceID="xmp.iid:4F18EC5019EEE11193D8B71EAE518AB8"
         stRef:documentID="xmp.did:4C18EC5019EEE11193D8B71EAE518AB8"
         stRef:fromPart="time:0d75776f25600"
         stRef:toPart="time:0d75776f25600"
         stRef:maskMarkers="All"/>
       </rdf:Bag>
      </xmpMM:Ingredients>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       xmpMM:OriginalDocumentID="xmp.did:0880117407206811AFFDE2CE2A0925F6"
       xmpMM:DocumentID="xmp.did:0880117407206811AFFDE2CE2A0925F6"
       xmpMM:InstanceID="xmp.iid:1FFC3CFF4DEDE1118EE487FA45ED9BCA"
       xmp:CreatorTool="After Effects"
       xmp:MetadataDate="2012-08-23T20:13:58+02:00"
       xmp:CreateDate="2009-01-22T21:30:55Z"
       xmp:ModifyDate="2012-08-23T18:13:48Z"
       xmpDM:videoPixelAspectRatio="768/702"
       xmpDM:videoFieldOrder="Progressive"
       xmpDM:videoAlphaMode="none"
       xmpDM:videoFrameRate="30.000000">
      <xmpMM:History>
       <rdf:Seq>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:0980117407206811AFFDE2CE2A0925F6"
         stEvt:when="2010-08-22T17:14:26-05:00"
         stEvt:softwareAgent="Adobe After Effects 9.0"
         stEvt:changed="/metadata"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:A6D0CA5C09206811AFFDE2CE2A0925F6"
         stEvt:when="2010-08-22T17:31:46-05:00"
         stEvt:softwareAgent="Adobe After Effects 9.0"
         stEvt:changed="/metadata"/>
        <rdf:li
         stEvt:action="modified"
         stEvt:parameters="unknown modifications"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:0E460B79130BE011ACA4902236DF1735"
         stEvt:when="2010-12-18T21:30:30-05:00"
         stEvt:softwareAgent="Adobe After Effects CS5 10.0 Windows 64"
         stEvt:changed="/metadata;/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:11460B79130BE011ACA4902236DF1735"
         stEvt:when="2010-12-18T21:31-05:00"
         stEvt:softwareAgent="Adobe After Effects CS5 10.0 Windows 64"
         stEvt:changed="/metadata"/>
        <rdf:li
         stEvt:action="modified"
         stEvt:parameters="unknown modifications"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:1EFC3CFF4DEDE1118EE487FA45ED9BCA"
         stEvt:when="2012-08-23T20:13:48+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/metadata;/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:1FFC3CFF4DEDE1118EE487FA45ED9BCA"
         stEvt:when="2012-08-23T20:13:58+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/metadata"/>
       </rdf:Seq>
      </xmpMM:History>
      <xmpDM:videoFrameSize
       stDim:w="720"
       stDim:h="576"
       stDim:unit="pixel"/>
      <xmpDM:duration
       xmpDM:value="1600"
       xmpDM:scale="1/600"/>
      <creatorAtom:windowsAtom
       creatorAtom:extension=".aep"
       creatorAtom:invocationFlags="-ep"/>
      <creatorAtom:macAtom
       creatorAtom:applicationCode="1180193859"
       creatorAtom:invocationAppleEvent="1131559026"/>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       xmp:CreateDate="2012-09-01T17:36:31Z"
       xmp:ModifyDate="2012-09-01T16:39:22Z"
       xmp:CreatorTool="Adobe After Effects CS5.5 10.5 "
       xmp:MetadataDate="2012-09-01T18:39:32+02:00"
       xmpDM:startTimeScale="25"
       xmpDM:startTimeSampleSize="1"
       xmpDM:videoFrameRate="25.000000"
       xmpDM:videoFieldOrder="Progressive"
       xmpDM:videoPixelAspectRatio="1/1"
       xmpDM:videoAlphaMode="none"
       dc:format="QuickTime"
       xmpMM:InstanceID="xmp.iid:2C0FA19553F4E1118E5EF593D63F81DA"
       xmpMM:DocumentID="xmp.did:8254602F53F4E11188E1CBA262DF7ED4"
       xmpMM:OriginalDocumentID="xmp.did:5F294BA351EDE111BA25D0E0299F7394">
      <xmpDM:altTimecode
       xmpDM:timeValue="00:00:00:00"
       xmpDM:timeFormat="25Timecode"/>
      <xmpDM:videoFrameSize
       stDim:w="1920"
       stDim:h="1080"
       stDim:unit="pixel"/>
      <xmpDM:startTimecode
       xmpDM:timeFormat="25Timecode"
       xmpDM:timeValue="00:00:00:00"/>
      <xmpDM:duration
       xmpDM:value="1200"
       xmpDM:scale="1/25"/>
      <xmpMM:History>
       <rdf:Seq>
        <rdf:li
         stEvt:action="created"
         stEvt:instanceID="xmp.iid:5F294BA351EDE111BA25D0E0299F7394"
         stEvt:when="2012-08-23T20:37:48+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:413A8F1652EDE111BA25D0E0299F7394"
         stEvt:when="2012-08-23T21:01:17+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:423A8F1652EDE111BA25D0E0299F7394"
         stEvt:when="2012-08-23T21:16:52+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:433A8F1652EDE111BA25D0E0299F7394"
         stEvt:when="2012-08-23T21:18+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:C3EFDF4057EDE111BA25D0E0299F7394"
         stEvt:when="2012-08-23T21:24:08+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:FF4E80EA19EEE11180F5C2DBA05DF1FB"
         stEvt:when="2012-08-24T21:21:46+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:004F80EA19EEE11180F5C2DBA05DF1FB"
         stEvt:when="2012-08-24T21:32:46+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:014F80EA19EEE11180F5C2DBA05DF1FB"
         stEvt:when="2012-08-24T21:34:14+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:1944063158F0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T17:06:02+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:1A44063158F0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T17:09:50+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:1B44063158F0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T17:13:27+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:1C44063158F0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T17:14:15+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:1D44063158F0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T17:56:21+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:1E44063158F0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T17:56:38+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:1F44063158F0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T17:59:26+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:B55311F762F0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T18:19:24+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:B65311F762F0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T18:20:05+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:B75311F762F0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T18:39:36+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:B85311F762F0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T18:42:11+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:B95311F762F0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T18:54:21+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:BA5311F762F0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T19:00:19+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:BB5311F762F0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T19:07+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:BC5311F762F0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T19:13:28+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:BD5311F762F0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T19:22:26+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:BE5311F762F0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T19:23:21+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:BF5311F762F0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T19:25:57+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:E8DAC1686CF0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T19:27+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:E9DAC1686CF0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T19:46:02+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:EADAC1686CF0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T19:48:10+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:EBDAC1686CF0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T19:49:42+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:ECDAC1686CF0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T19:54:28+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:EDDAC1686CF0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T19:58:22+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:EEDAC1686CF0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T20:02:50+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:EFDAC1686CF0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T20:06:16+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:F0DAC1686CF0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T20:18:58+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:F1DAC1686CF0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T20:30:10+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:F2DAC1686CF0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T20:31:55+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:47B98B7A75F0E111A00BD719F39AD641"
         stEvt:when="2012-08-27T20:49:30+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:DC8D9742FBF1E1118CF6B6FF225B6458"
         stEvt:when="2012-08-29T19:07:07+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:E28D9742FBF1E1118CF6B6FF225B6458"
         stEvt:when="2012-08-29T19:08:02+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:144767901FF4E111AF3ABB745A442EB4"
         stEvt:when="2012-09-01T12:47:26+02:00"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:154767901FF4E111AF3ABB745A442EB4"
         stEvt:when="2012-09-01T12:47:43+02:00"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:164767901FF4E111AF3ABB745A442EB4"
         stEvt:when="2012-09-01T13:06:02+02:00"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:9409786E4FF4E1118B20F7D2ADD7C03B"
         stEvt:when="2012-09-01T18:16:52+02:00"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:4D9DBE7952F4E11188E1CBA262DF7ED4"
         stEvt:when="2012-09-01T18:33:56+02:00"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:4E9DBE7952F4E11188E1CBA262DF7ED4"
         stEvt:when="2012-09-01T18:36:01+02:00"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:4F9DBE7952F4E11188E1CBA262DF7ED4"
         stEvt:when="2012-09-01T18:36:31+02:00"
         stEvt:changed="/content"/>
        <rdf:li
         stEvt:action="converted"
         stEvt:parameters="from application/vnd.adobe.aftereffects.project to video/mov"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:8254602F53F4E11188E1CBA262DF7ED4"
         stEvt:when="2012-09-01T18:36:31+02:00"
         stEvt:changed="/;/metadata"/>
        <rdf:li
         stEvt:action="modified"
         stEvt:parameters="unknown modifications"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:2B0FA19553F4E1118E5EF593D63F81DA"
         stEvt:when="2012-09-01T18:39:22+02:00"
         stEvt:softwareAgent="Adobe Premiere Pro 5.5"
         stEvt:changed="/metadata;/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:2C0FA19553F4E1118E5EF593D63F81DA"
         stEvt:when="2012-09-01T18:39:32+02:00"
         stEvt:softwareAgent="Adobe Premiere Pro 5.5"
         stEvt:changed="/metadata"/>
       </rdf:Seq>
      </xmpMM:History>
      <xmpMM:Ingredients>
       <rdf:Bag>
        <rdf:li
         stRef:instanceID="xmp.iid:8154602F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d1228800f25600"
         stRef:toPart="time:0d1228800f25600"
         stRef:maskMarkers="None"/>
       </rdf:Bag>
      </xmpMM:Ingredients>
      <xmpMM:DerivedFrom
       stRef:instanceID="xmp.iid:4F9DBE7952F4E11188E1CBA262DF7ED4"
       stRef:documentID="xmp.did:5F294BA351EDE111BA25D0E0299F7394"
       stRef:originalDocumentID="xmp.did:5F294BA351EDE111BA25D0E0299F7394"/>
      <creatorAtom:windowsAtom
       creatorAtom:extension=".aep"
       creatorAtom:invocationFlags="-ep"/>
      <creatorAtom:macAtom
       creatorAtom:applicationCode="1180193859"
       creatorAtom:invocationAppleEvent="1131559026"/>
      <creatorAtom:aeProjectLink
       creatorAtom:renderTimeStamp="24222"
       creatorAtom:compositionID="2"
       creatorAtom:renderQueueItemID="190"
       creatorAtom:renderOutputModuleIndex="0"
       creatorAtom:fullPath="E:\alyaspot\Posnetki\Nova mapa\burn_efekt_trailer.aep"/>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       xmp:CreateDate="2012-08-12T16:42:59Z"
       xmp:ModifyDate="2012-08-19T10:28:04Z"
       xmp:MetadataDate="2012-08-19T12:28:12+02:00"
       xmpDM:videoAlphaMode="none"
       xmpDM:audioSampleRate="48000"
       xmpDM:audioSampleType="Compressed"
       xmpDM:audioChannelType="Stereo"
       xmpDM:videoFrameRate="25.000000"
       xmpMM:InstanceID="xmp.iid:3C5F4D8EE8E9E111B624CD95E5AA602F"
       xmpMM:DocumentID="xmp.did:395F4D8EE8E9E111B624CD95E5AA602F"
       xmpMM:OriginalDocumentID="xmp.did:395F4D8EE8E9E111B624CD95E5AA602F">
      <xmpDM:videoFrameSize
       stDim:w="1920"
       stDim:h="1080"
       stDim:unit="pixel"/>
      <xmpDM:duration
       xmpDM:value="2445000"
       xmpDM:scale="1/25000"/>
      <xmpMM:History>
       <rdf:Seq>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:3A5F4D8EE8E9E111B624CD95E5AA602F"
         stEvt:when="2012-08-19T12:28:04+02:00"
         stEvt:softwareAgent="Adobe Premiere Pro 5.5"
         stEvt:changed="/metadata;/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:3C5F4D8EE8E9E111B624CD95E5AA602F"
         stEvt:when="2012-08-19T12:28:12+02:00"
         stEvt:softwareAgent="Adobe Premiere Pro 5.5"
         stEvt:changed="/metadata"/>
       </rdf:Seq>
      </xmpMM:History>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       xmp:CreateDate="2012-08-12T14:35:35Z"
       xmp:ModifyDate="2012-08-19T10:28:02Z"
       xmp:MetadataDate="2012-08-19T12:28:22+02:00"
       xmpDM:videoAlphaMode="none"
       xmpDM:audioSampleRate="48000"
       xmpDM:audioSampleType="Compressed"
       xmpDM:audioChannelType="Stereo"
       xmpDM:videoFrameRate="25.000000"
       xmpMM:InstanceID="xmp.iid:3D5F4D8EE8E9E111B624CD95E5AA602F"
       xmpMM:DocumentID="xmp.did:355F4D8EE8E9E111B624CD95E5AA602F"
       xmpMM:OriginalDocumentID="xmp.did:355F4D8EE8E9E111B624CD95E5AA602F">
      <xmpDM:videoFrameSize
       stDim:w="1920"
       stDim:h="1080"
       stDim:unit="pixel"/>
      <xmpDM:duration
       xmpDM:value="2555000"
       xmpDM:scale="1/25000"/>
      <xmpMM:History>
       <rdf:Seq>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:365F4D8EE8E9E111B624CD95E5AA602F"
         stEvt:when="2012-08-19T12:28:02+02:00"
         stEvt:softwareAgent="Adobe Premiere Pro 5.5"
         stEvt:changed="/metadata;/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:3D5F4D8EE8E9E111B624CD95E5AA602F"
         stEvt:when="2012-08-19T12:28:22+02:00"
         stEvt:softwareAgent="Adobe Premiere Pro 5.5"
         stEvt:changed="/metadata"/>
       </rdf:Seq>
      </xmpMM:History>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       xmp:CreatorTool="Adobe Photoshop CS5 Windows"
       xmp:CreateDate="2010-11-26T18:28:24+01:00"
       xmp:MetadataDate="2011-03-15T22:32:25+01:00"
       xmp:ModifyDate="2011-03-15T22:32:15+01:00"
       dc:format="application/vnd.adobe.photoshop"
       xmpMM:InstanceID="xmp.iid:44D1D8A6474FE0119FC1D4121F99B6A6"
       xmpMM:DocumentID="xmp.did:551E949282F9DF11A7BAFCB45617540F"
       xmpMM:OriginalDocumentID="xmp.did:551E949282F9DF11A7BAFCB45617540F"
       photoshop:ColorMode="3"
       xmpDM:videoPixelAspectRatio="10000/10000"
       xmpDM:videoFieldOrder="Progressive"
       xmpDM:videoAlphaMode="straight"
       xmpDM:videoFrameRate="0.000000"
       tiff:Orientation="1"
       tiff:XResolution="720000/10000"
       tiff:YResolution="720000/10000"
       tiff:ResolutionUnit="2"
       exif:ColorSpace="65535"
       exif:PixelXDimension="1920"
       exif:PixelYDimension="1080">
      <xmpMM:History>
       <rdf:Seq>
        <rdf:li
         stEvt:action="created"
         stEvt:instanceID="xmp.iid:551E949282F9DF11A7BAFCB45617540F"
         stEvt:when="2010-11-26T18:28:24+01:00"
         stEvt:softwareAgent="Adobe Photoshop CS5 Windows"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:B29F167B1FF9DF11AD2EC76BC855A93E"
         stEvt:when="2010-11-26T18:28:38+01:00"
         stEvt:softwareAgent="Adobe Premiere Pro 5.0"
         stEvt:changed="/metadata"/>
        <rdf:li
         stEvt:action="modified"
         stEvt:parameters="unknown modifications"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:43D1D8A6474FE0119FC1D4121F99B6A6"
         stEvt:when="2011-03-15T22:32:15+01:00"
         stEvt:softwareAgent="Adobe Premiere Pro 5.0"
         stEvt:changed="/metadata;/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:44D1D8A6474FE0119FC1D4121F99B6A6"
         stEvt:when="2011-03-15T22:32:25+01:00"
         stEvt:softwareAgent="Adobe Premiere Pro 5.0"
         stEvt:changed="/metadata"/>
       </rdf:Seq>
      </xmpMM:History>
      <xmpDM:videoFrameSize
       stDim:w="1920"
       stDim:h="1080"
       stDim:unit="pixel"/>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       xmp:CreateDate="2012-08-12T15:36:08Z"
       xmp:ModifyDate="2012-08-24T16:12:13Z"
       xmp:MetadataDate="2012-08-24T18:12:23+02:00"
       xmpDM:videoAlphaMode="none"
       xmpDM:audioSampleRate="48000"
       xmpDM:audioSampleType="Compressed"
       xmpDM:audioChannelType="Stereo"
       xmpDM:videoFrameRate="25.000000"
       xmpMM:InstanceID="xmp.iid:45F9717706EEE1119BFCD6521D1187C9"
       xmpMM:DocumentID="xmp.did:43F9717706EEE1119BFCD6521D1187C9"
       xmpMM:OriginalDocumentID="xmp.did:43F9717706EEE1119BFCD6521D1187C9">
      <xmpDM:videoFrameSize
       stDim:w="1920"
       stDim:h="1080"
       stDim:unit="pixel"/>
      <xmpDM:duration
       xmpDM:value="3316000"
       xmpDM:scale="1/25000"/>
      <xmpMM:History>
       <rdf:Seq>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:44F9717706EEE1119BFCD6521D1187C9"
         stEvt:when="2012-08-24T18:12:13+02:00"
         stEvt:softwareAgent="Adobe Premiere Pro 5.5"
         stEvt:changed="/metadata;/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:45F9717706EEE1119BFCD6521D1187C9"
         stEvt:when="2012-08-24T18:12:23+02:00"
         stEvt:softwareAgent="Adobe Premiere Pro 5.5"
         stEvt:changed="/metadata"/>
       </rdf:Seq>
      </xmpMM:History>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       xmp:CreatorTool="Adobe Photoshop CS5.1 Windows"
       xmp:CreateDate="2012-08-24T20:27:09+02:00"
       xmp:MetadataDate="2012-08-24T20:27:21+02:00"
       xmp:ModifyDate="2012-08-24T20:27:21+02:00"
       photoshop:ColorMode="3"
       photoshop:ICCProfile="sRGB IEC61966-2.1"
       dc:format="image/jpeg"
       xmpMM:InstanceID="xmp.iid:4F18EC5019EEE11193D8B71EAE518AB8"
       xmpMM:DocumentID="xmp.did:4C18EC5019EEE11193D8B71EAE518AB8"
       xmpMM:OriginalDocumentID="xmp.did:4C18EC5019EEE11193D8B71EAE518AB8"
       tiff:Orientation="1"
       tiff:XResolution="720000/10000"
       tiff:YResolution="720000/10000"
       tiff:ResolutionUnit="2"
       exif:ColorSpace="1"
       exif:PixelXDimension="1920"
       exif:PixelYDimension="1080">
      <xmpMM:History>
       <rdf:Seq>
        <rdf:li
         stEvt:action="created"
         stEvt:instanceID="xmp.iid:4C18EC5019EEE11193D8B71EAE518AB8"
         stEvt:when="2012-08-24T20:27:09+02:00"
         stEvt:softwareAgent="Adobe Photoshop CS5.1 Windows"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:4D18EC5019EEE11193D8B71EAE518AB8"
         stEvt:when="2012-08-24T20:27:12+02:00"
         stEvt:softwareAgent="Adobe Photoshop CS5.1 Windows"
         stEvt:changed="/"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:4E18EC5019EEE11193D8B71EAE518AB8"
         stEvt:when="2012-08-24T20:27:21+02:00"
         stEvt:softwareAgent="Adobe Photoshop CS5.1 Windows"
         stEvt:changed="/"/>
        <rdf:li
         stEvt:action="converted"
         stEvt:parameters="from application/vnd.adobe.photoshop to image/jpeg"/>
        <rdf:li
         stEvt:action="derived"
         stEvt:parameters="converted from application/vnd.adobe.photoshop to image/jpeg"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:4F18EC5019EEE11193D8B71EAE518AB8"
         stEvt:when="2012-08-24T20:27:21+02:00"
         stEvt:softwareAgent="Adobe Photoshop CS5.1 Windows"
         stEvt:changed="/"/>
       </rdf:Seq>
      </xmpMM:History>
      <xmpMM:DerivedFrom
       stRef:instanceID="xmp.iid:4E18EC5019EEE11193D8B71EAE518AB8"
       stRef:documentID="xmp.did:4C18EC5019EEE11193D8B71EAE518AB8"
       stRef:originalDocumentID="xmp.did:4C18EC5019EEE11193D8B71EAE518AB8"/>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:509DBE7952F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">240mask1080.psd</rdf:li>
       </rdf:Alt>
      </dc:title>
      <xmpMM:Ingredients>
       <rdf:Bag>
        <rdf:li
         stRef:instanceID="xmp.iid:44D1D8A6474FE0119FC1D4121F99B6A6"
         stRef:documentID="xmp.did:551E949282F9DF11A7BAFCB45617540F"
         stRef:fromPart="time:0d1228800f25600"
         stRef:toPart="time:0d1228800f25600"
         stRef:maskMarkers="All"/>
       </rdf:Bag>
      </xmpMM:Ingredients>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:519DBE7952F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">Black Solid 4</rdf:li>
       </rdf:Alt>
      </dc:title>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       xmp:CreateDate="2012-08-12T15:54:13Z"
       xmp:ModifyDate="2012-08-22T16:00:38Z"
       xmp:MetadataDate="2012-08-22T18:00:48+02:00"
       xmpDM:videoAlphaMode="none"
       xmpDM:audioSampleRate="48000"
       xmpDM:audioSampleType="Compressed"
       xmpDM:audioChannelType="Stereo"
       xmpDM:videoFrameRate="25.000000"
       xmpMM:InstanceID="xmp.iid:5255F98372ECE111B9A38987502929D1"
       xmpMM:DocumentID="xmp.did:5055F98372ECE111B9A38987502929D1"
       xmpMM:OriginalDocumentID="xmp.did:5055F98372ECE111B9A38987502929D1">
      <xmpDM:videoFrameSize
       stDim:w="1920"
       stDim:h="1080"
       stDim:unit="pixel"/>
      <xmpDM:duration
       xmpDM:value="2293000"
       xmpDM:scale="1/25000"/>
      <xmpMM:History>
       <rdf:Seq>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:5155F98372ECE111B9A38987502929D1"
         stEvt:when="2012-08-22T18:00:38+02:00"
         stEvt:softwareAgent="Adobe Premiere Pro 5.5"
         stEvt:changed="/metadata;/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:5255F98372ECE111B9A38987502929D1"
         stEvt:when="2012-08-22T18:00:48+02:00"
         stEvt:softwareAgent="Adobe Premiere Pro 5.5"
         stEvt:changed="/metadata"/>
       </rdf:Seq>
      </xmpMM:History>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:529DBE7952F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">Light Gray Solid 1</rdf:li>
       </rdf:Alt>
      </dc:title>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:539DBE7952F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">Black Solid 3</rdf:li>
       </rdf:Alt>
      </dc:title>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       xmp:CreateDate="2012-08-12T15:58Z"
       xmp:ModifyDate="2012-08-22T16:07:25Z"
       xmp:MetadataDate="2012-08-22T18:07:35+02:00"
       xmpDM:videoAlphaMode="none"
       xmpDM:audioSampleRate="48000"
       xmpDM:audioSampleType="Compressed"
       xmpDM:audioChannelType="Stereo"
       xmpDM:videoFrameRate="25.000000"
       xmpMM:InstanceID="xmp.iid:5555F98372ECE111B9A38987502929D1"
       xmpMM:DocumentID="xmp.did:5355F98372ECE111B9A38987502929D1"
       xmpMM:OriginalDocumentID="xmp.did:5355F98372ECE111B9A38987502929D1">
      <xmpDM:videoFrameSize
       stDim:w="1920"
       stDim:h="1080"
       stDim:unit="pixel"/>
      <xmpDM:duration
       xmpDM:value="3270000"
       xmpDM:scale="1/25000"/>
      <xmpMM:History>
       <rdf:Seq>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:5455F98372ECE111B9A38987502929D1"
         stEvt:when="2012-08-22T18:07:25+02:00"
         stEvt:softwareAgent="Adobe Premiere Pro 5.5"
         stEvt:changed="/metadata;/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:5555F98372ECE111B9A38987502929D1"
         stEvt:when="2012-08-22T18:07:35+02:00"
         stEvt:softwareAgent="Adobe Premiere Pro 5.5"
         stEvt:changed="/metadata"/>
       </rdf:Seq>
      </xmpMM:History>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.comp"
       xmpMM:InstanceID="xmp.iid:55955D2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">Comp 1</rdf:li>
       </rdf:Alt>
      </dc:title>
      <xmpMM:Ingredients>
       <rdf:Bag>
        <rdf:li
         stRef:instanceID="xmp.iid:1CF95C2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d75776f25600"
         stRef:toPart="time:0d75776f25600"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:1DF95C2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d75776f25600"
         stRef:toPart="time:0d75776f25600"
         stRef:maskMarkers="None"/>
       </rdf:Bag>
      </xmpMM:Ingredients>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:56955D2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">Comp 1</rdf:li>
       </rdf:Alt>
      </dc:title>
      <xmpMM:Ingredients>
       <rdf:Bag>
        <rdf:li
         stRef:instanceID="xmp.iid:55955D2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d75776f25600"
         stRef:toPart="time:0d75776f25600"
         stRef:maskMarkers="None"/>
       </rdf:Bag>
      </xmpMM:Ingredients>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:57955D2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">film burn 8-Apple ProRes 422 (LT).mov</rdf:li>
       </rdf:Alt>
      </dc:title>
      <xmpMM:Ingredients>
       <rdf:Bag>
        <rdf:li
         stRef:instanceID="xmp.iid:AC28A10A52EDE111BA25D0E0299F7394"
         stRef:documentID="xmp.did:3A8FC90452EDE111BA25D0E0299F7394"
         stRef:fromPart="time:26368f25600d39936f25600"
         stRef:toPart="time:26368f25600d39936f25600"
         stRef:maskMarkers="All"/>
       </rdf:Bag>
      </xmpMM:Ingredients>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:58955D2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">film burn 1-Apple ProRes 422 (LT).mov</rdf:li>
       </rdf:Alt>
      </dc:title>
      <xmpMM:Ingredients>
       <rdf:Bag>
        <rdf:li
         stRef:instanceID="xmp.iid:C2F75C0552EDE111BA25D0E0299F7394"
         stRef:documentID="xmp.did:64294BA351EDE111BA25D0E0299F7394"
         stRef:fromPart="time:0d660441600000f254016000000"
         stRef:toPart="time:0d660441600000f254016000000"
         stRef:maskMarkers="All"/>
       </rdf:Bag>
      </xmpMM:Ingredients>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       xmp:CreatorTool="Adobe Photoshop CS5.1 Windows"
       xmp:CreateDate="2012-08-27T18:27:20+02:00"
       xmp:MetadataDate="2012-08-27T18:27:20+02:00"
       xmp:ModifyDate="2012-08-27T18:27:20+02:00"
       xmpMM:InstanceID="xmp.iid:5933D21264F0E11180B8D0D5AB388124"
       xmpMM:DocumentID="xmp.did:5833D21264F0E11180B8D0D5AB388124"
       xmpMM:OriginalDocumentID="xmp.did:5833D21264F0E11180B8D0D5AB388124"
       dc:format="image/jpeg"
       photoshop:ColorMode="3"
       photoshop:ICCProfile="sRGB IEC61966-2.1"
       tiff:Orientation="1"
       tiff:XResolution="720000/10000"
       tiff:YResolution="720000/10000"
       tiff:ResolutionUnit="2"
       exif:ColorSpace="1"
       exif:PixelXDimension="1920"
       exif:PixelYDimension="1080">
      <xmpMM:History>
       <rdf:Seq>
        <rdf:li
         stEvt:action="created"
         stEvt:instanceID="xmp.iid:5833D21264F0E11180B8D0D5AB388124"
         stEvt:when="2012-08-27T18:27:20+02:00"
         stEvt:softwareAgent="Adobe Photoshop CS5.1 Windows"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:5933D21264F0E11180B8D0D5AB388124"
         stEvt:when="2012-08-27T18:27:20+02:00"
         stEvt:softwareAgent="Adobe Photoshop CS5.1 Windows"
         stEvt:changed="/"/>
       </rdf:Seq>
      </xmpMM:History>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:59955D2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">film burn 17-Apple ProRes 422 (LT).mov</rdf:li>
       </rdf:Alt>
      </dc:title>
      <xmpMM:Ingredients>
       <rdf:Bag>
        <rdf:li
         stRef:instanceID="xmp.iid:B4DE9E0A52EDE111BA25D0E0299F7394"
         stRef:documentID="xmp.did:DE63290552EDE111BA25D0E0299F7394"
         stRef:fromPart="time:0d753580800000f254016000000"
         stRef:toPart="time:0d753580800000f254016000000"
         stRef:maskMarkers="All"/>
       </rdf:Bag>
      </xmpMM:Ingredients>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:5A955D2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">Real filmburn.mov</rdf:li>
       </rdf:Alt>
      </dc:title>
      <xmpMM:Ingredients>
       <rdf:Bag>
        <rdf:li
         stRef:instanceID="xmp.iid:1FFC3CFF4DEDE1118EE487FA45ED9BCA"
         stRef:documentID="xmp.did:0880117407206811AFFDE2CE2A0925F6"
         stRef:fromPart="time:0d677376000000f254016000000"
         stRef:toPart="time:0d677376000000f254016000000"
         stRef:maskMarkers="All"/>
       </rdf:Bag>
      </xmpMM:Ingredients>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:5B955D2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">Real filmburn.mov</rdf:li>
       </rdf:Alt>
      </dc:title>
      <xmpMM:Ingredients>
       <rdf:Bag>
        <rdf:li
         stRef:instanceID="xmp.iid:1FFC3CFF4DEDE1118EE487FA45ED9BCA"
         stRef:documentID="xmp.did:0880117407206811AFFDE2CE2A0925F6"
         stRef:fromPart="time:0d677376000000f254016000000"
         stRef:toPart="time:0d677376000000f254016000000"
         stRef:maskMarkers="All"/>
       </rdf:Bag>
      </xmpMM:Ingredients>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:5C955D2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">Real filmburn.mov</rdf:li>
       </rdf:Alt>
      </dc:title>
      <xmpMM:Ingredients>
       <rdf:Bag>
        <rdf:li
         stRef:instanceID="xmp.iid:1FFC3CFF4DEDE1118EE487FA45ED9BCA"
         stRef:documentID="xmp.did:0880117407206811AFFDE2CE2A0925F6"
         stRef:fromPart="time:0d677376000000f254016000000"
         stRef:toPart="time:0d677376000000f254016000000"
         stRef:maskMarkers="All"/>
       </rdf:Bag>
      </xmpMM:Ingredients>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:5D955D2F53F4E11188E1CBA262DF7ED4"
       xmpMM:DocumentID="xmp.did:5D955D2F53F4E11188E1CBA262DF7ED4"
       xmpMM:OriginalDocumentID="xmp.did:5D955D2F53F4E11188E1CBA262DF7ED4"
       xmp:MetadataDate="2012-09-01T18:53:49+02:00">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">trailer_v1.mp4</rdf:li>
       </rdf:Alt>
      </dc:title>
      <xmpMM:Ingredients>
       <rdf:Bag>
        <rdf:li
         stRef:instanceID="xmp.iid:FE4E80EA19EEE11180F5C2DBA05DF1FB"
         stRef:documentID="xmp.did:D3BEB3351AEEE1119BFCD6521D1187C9"
         stRef:fromPart="time:944128f25600d157696f25600"
         stRef:toPart="time:944128f25600d157696f25600"
         stRef:maskMarkers="All"/>
       </rdf:Bag>
      </xmpMM:Ingredients>
      <xmpMM:History>
       <rdf:Seq>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:2F0FA19553F4E1118E5EF593D63F81DA"
         stEvt:when="2012-09-01T18:53:49+02:00"
         stEvt:softwareAgent="Adobe Premiere Pro 5.5"
         stEvt:changed="/metadata"/>
       </rdf:Seq>
      </xmpMM:History>
      <xmpDM:Tracks>
       <rdf:Bag>
        <rdf:li>
         <rdf:Description
          xmpDM:trackName="layer_markers"
          xmpDM:frameRate="f25600">
         <xmpDM:markers>
          <rdf:Seq>
           <rdf:li
            xmpDM:startTime="1031168"
            xmpDM:duration="1024"/>
          </rdf:Seq>
         </xmpDM:markers>
         </rdf:Description>
        </rdf:li>
       </rdf:Bag>
      </xmpDM:Tracks>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:5E955D2F53F4E11188E1CBA262DF7ED4"
       xmpMM:DocumentID="xmp.did:5E955D2F53F4E11188E1CBA262DF7ED4"
       xmpMM:OriginalDocumentID="xmp.did:5E955D2F53F4E11188E1CBA262DF7ED4"
       xmp:MetadataDate="2012-09-01T18:53:49+02:00">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">trailer_v1.mp4</rdf:li>
       </rdf:Alt>
      </dc:title>
      <xmpMM:Ingredients>
       <rdf:Bag>
        <rdf:li
         stRef:instanceID="xmp.iid:FE4E80EA19EEE11180F5C2DBA05DF1FB"
         stRef:documentID="xmp.did:D3BEB3351AEEE1119BFCD6521D1187C9"
         stRef:fromPart="time:819200f25600d124928f25600"
         stRef:toPart="time:819200f25600d124928f25600"
         stRef:maskMarkers="All"/>
       </rdf:Bag>
      </xmpMM:Ingredients>
      <xmpMM:History>
       <rdf:Seq>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:300FA19553F4E1118E5EF593D63F81DA"
         stEvt:when="2012-09-01T18:53:49+02:00"
         stEvt:softwareAgent="Adobe Premiere Pro 5.5"
         stEvt:changed="/metadata"/>
       </rdf:Seq>
      </xmpMM:History>
      <xmpDM:Tracks>
       <rdf:Bag>
        <rdf:li>
         <rdf:Description
          xmpDM:trackName="layer_markers"
          xmpDM:frameRate="f25600">
         <xmpDM:markers>
          <rdf:Seq>
           <rdf:li
            xmpDM:startTime="925696"
            xmpDM:duration="1024"/>
           <rdf:li
            xmpDM:startTime="940032"
            xmpDM:duration="1024"/>
          </rdf:Seq>
         </xmpDM:markers>
         </rdf:Description>
        </rdf:li>
       </rdf:Bag>
      </xmpDM:Tracks>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:5F955D2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">trailer_v1.mp4</rdf:li>
       </rdf:Alt>
      </dc:title>
      <xmpMM:Ingredients>
       <rdf:Bag>
        <rdf:li
         stRef:instanceID="xmp.iid:FE4E80EA19EEE11180F5C2DBA05DF1FB"
         stRef:documentID="xmp.did:D3BEB3351AEEE1119BFCD6521D1187C9"
         stRef:fromPart="time:684032f25600d135168f25600"
         stRef:toPart="time:684032f25600d135168f25600"
         stRef:maskMarkers="All"/>
       </rdf:Bag>
      </xmpMM:Ingredients>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       xmpMM:InstanceID="xmp.iid:661D67351FF4E111A01FA07C9D969101"
       xmpMM:DocumentID="xmp.did:641D67351FF4E111A01FA07C9D969101"
       xmpMM:OriginalDocumentID="xmp.did:641D67351FF4E111A01FA07C9D969101"
       xmp:MetadataDate="2012-09-01T12:24:37+02:00"
       xmp:ModifyDate="2012-09-01T12:24:27+02:00"
       xmpDM:audioSampleRate="44100"
       xmpDM:audioSampleType="16Int"
       xmpDM:audioChannelType="Stereo">
      <xmpMM:History>
       <rdf:Seq>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:651D67351FF4E111A01FA07C9D969101"
         stEvt:when="2012-09-01T12:24:27+02:00"
         stEvt:softwareAgent="Adobe Premiere Pro 5.5"
         stEvt:changed="/metadata;/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:661D67351FF4E111A01FA07C9D969101"
         stEvt:when="2012-09-01T12:24:37+02:00"
         stEvt:softwareAgent="Adobe Premiere Pro 5.5"
         stEvt:changed="/metadata"/>
       </rdf:Seq>
      </xmpMM:History>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:8054602F53F4E11188E1CBA262DF7ED4"
       xmpMM:DocumentID="xmp.did:8054602F53F4E11188E1CBA262DF7ED4"
       xmpMM:OriginalDocumentID="xmp.did:8054602F53F4E11188E1CBA262DF7ED4"
       xmp:MetadataDate="2012-09-01T18:53:49+02:00">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">trailer_v1.mp4</rdf:li>
       </rdf:Alt>
      </dc:title>
      <xmpMM:Ingredients>
       <rdf:Bag>
        <rdf:li
         stRef:instanceID="xmp.iid:FE4E80EA19EEE11180F5C2DBA05DF1FB"
         stRef:documentID="xmp.did:D3BEB3351AEEE1119BFCD6521D1187C9"
         stRef:fromPart="time:0d1101824f25600"
         stRef:toPart="time:0d1101824f25600"
         stRef:maskMarkers="All"/>
       </rdf:Bag>
      </xmpMM:Ingredients>
      <xmpMM:History>
       <rdf:Seq>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:310FA19553F4E1118E5EF593D63F81DA"
         stEvt:when="2012-09-01T18:53:49+02:00"
         stEvt:softwareAgent="Adobe Premiere Pro 5.5"
         stEvt:changed="/metadata"/>
       </rdf:Seq>
      </xmpMM:History>
      <xmpDM:Tracks>
       <rdf:Bag>
        <rdf:li>
         <rdf:Description
          xmpDM:trackName="layer_markers"
          xmpDM:frameRate="f25600">
         <xmpDM:markers>
          <rdf:Seq>
           <rdf:li
            xmpDM:startTime="302080"
            xmpDM:duration="1024"/>
           <rdf:li
            xmpDM:startTime="925696"
            xmpDM:duration="1024"/>
           <rdf:li
            xmpDM:startTime="940032"
            xmpDM:duration="1024"/>
           <rdf:li
            xmpDM:startTime="1031168"
            xmpDM:duration="1024"/>
          </rdf:Seq>
         </xmpDM:markers>
         </rdf:Description>
        </rdf:li>
       </rdf:Bag>
      </xmpDM:Tracks>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.comp"
       xmpMM:InstanceID="xmp.iid:8154602F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">sekvenca_za_film_burn</rdf:li>
       </rdf:Alt>
      </dc:title>
      <xmpMM:Ingredients>
       <rdf:Bag>
        <rdf:li
         stRef:instanceID="xmp.iid:14F95C2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d76800f25600"
         stRef:toPart="time:461824f25600d76800f25600"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:1BF95C2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d72704f25600"
         stRef:toPart="time:0d72704f25600"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:509DBE7952F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d1228800f25600"
         stRef:toPart="time:0d1228800f25600"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:56955D2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d75776f25600"
         stRef:toPart="time:461824f25600d75776f25600"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:57955D2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:26368f25600d39936f25600"
         stRef:toPart="time:801792f25600d39936f25600"
         stRef:partMapping="linear-reversed"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:58955D2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d660441600000f254016000000"
         stRef:toPart="time:893952f25600d93184f25600"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:59955D2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d753580800000f254016000000"
         stRef:toPart="time:112640f25600d753580800000f254016000000"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:5A955D2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d677376000000f254016000000"
         stRef:toPart="time:265216f25600d677376000000f254016000000"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:5B955D2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d677376000000f254016000000"
         stRef:toPart="time:626688f25600d677376000000f254016000000"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:5C955D2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d677376000000f254016000000"
         stRef:toPart="time:69632f25600d677376000000f254016000000"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:5D955D2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:944128f25600d157696f25600"
         stRef:toPart="time:944128f25600d157696f25600"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:5E955D2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:819200f25600d124928f25600"
         stRef:toPart="time:819200f25600d124928f25600"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:5F955D2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:684032f25600d135168f25600"
         stRef:toPart="time:684032f25600d135168f25600"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:8054602F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d1101824f25600"
         stRef:toPart="time:0d1101824f25600"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:C5355C2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d128000f25600"
         stRef:toPart="time:1055744f25600d128000f25600"
         stRef:maskMarkers="None"/>
       </rdf:Bag>
      </xmpMM:Ingredients>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       xmp:CreateDate="2011-03-01T01:48:52Z"
       xmp:ModifyDate="2012-08-23T18:40:32Z"
       xmp:MetadataDate="2012-08-23T20:40:41+02:00"
       xmpDM:startTimeScale="3000"
       xmpDM:startTimeSampleSize="100"
       xmpDM:videoPixelAspectRatio="40/33"
       xmpDM:videoFieldOrder="Progressive"
       xmpDM:videoAlphaMode="none"
       xmpDM:videoFrameRate="30.000000"
       xmpMM:InstanceID="xmp.iid:AC28A10A52EDE111BA25D0E0299F7394"
       xmpMM:DocumentID="xmp.did:3A8FC90452EDE111BA25D0E0299F7394"
       xmpMM:OriginalDocumentID="xmp.did:3A8FC90452EDE111BA25D0E0299F7394">
      <xmpDM:altTimecode
       xmpDM:timeValue="00:00:00:00"
       xmpDM:timeFormat="30Timecode"/>
      <xmpDM:duration
       xmpDM:value="9700"
       xmpDM:scale="1/3000"/>
      <xmpDM:videoFrameSize
       stDim:w="720"
       stDim:h="480"
       stDim:unit="pixel"/>
      <xmpMM:History>
       <rdf:Seq>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:3B8FC90452EDE111BA25D0E0299F7394"
         stEvt:when="2012-08-23T20:40:32+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/metadata;/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:AC28A10A52EDE111BA25D0E0299F7394"
         stEvt:when="2012-08-23T20:40:41+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/metadata"/>
       </rdf:Seq>
      </xmpMM:History>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       xmp:CreateDate="2011-03-01T01:49:12Z"
       xmp:ModifyDate="2012-08-23T18:40:32Z"
       xmp:MetadataDate="2012-08-23T20:40:41+02:00"
       xmpDM:startTimeScale="3000"
       xmpDM:startTimeSampleSize="100"
       xmpDM:videoPixelAspectRatio="40/33"
       xmpDM:videoFieldOrder="Progressive"
       xmpDM:videoAlphaMode="none"
       xmpDM:videoFrameRate="30.000000"
       xmpMM:InstanceID="xmp.iid:B4DE9E0A52EDE111BA25D0E0299F7394"
       xmpMM:DocumentID="xmp.did:DE63290552EDE111BA25D0E0299F7394"
       xmpMM:OriginalDocumentID="xmp.did:DE63290552EDE111BA25D0E0299F7394">
      <xmpDM:altTimecode
       xmpDM:timeValue="00:00:00:00"
       xmpDM:timeFormat="30Timecode"/>
      <xmpDM:duration
       xmpDM:value="8900"
       xmpDM:scale="1/3000"/>
      <xmpDM:videoFrameSize
       stDim:w="720"
       stDim:h="480"
       stDim:unit="pixel"/>
      <xmpMM:History>
       <rdf:Seq>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:DF63290552EDE111BA25D0E0299F7394"
         stEvt:when="2012-08-23T20:40:32+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/metadata;/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:B4DE9E0A52EDE111BA25D0E0299F7394"
         stEvt:when="2012-08-23T20:40:41+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/metadata"/>
       </rdf:Seq>
      </xmpMM:History>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:C0355C2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">www.alya.info www.facebook.com/AlyaMusic www.facebook.com/QViewProduction</rdf:li>
       </rdf:Alt>
      </dc:title>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:C1355C2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">Coming soon</rdf:li>
       </rdf:Alt>
      </dc:title>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:C2355C2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">White Solid 2</rdf:li>
       </rdf:Alt>
      </dc:title>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       xmp:CreateDate="2011-03-01T01:48:34Z"
       xmp:ModifyDate="2012-08-23T18:40:31Z"
       xmp:MetadataDate="2012-08-23T20:40:41+02:00"
       xmpDM:startTimeScale="2997"
       xmpDM:startTimeSampleSize="100"
       xmpDM:videoPixelAspectRatio="40/33"
       xmpDM:videoFieldOrder="Progressive"
       xmpDM:videoAlphaMode="none"
       xmpDM:videoFrameRate="29.970030"
       xmpMM:InstanceID="xmp.iid:C2F75C0552EDE111BA25D0E0299F7394"
       xmpMM:DocumentID="xmp.did:64294BA351EDE111BA25D0E0299F7394"
       xmpMM:OriginalDocumentID="xmp.did:64294BA351EDE111BA25D0E0299F7394">
      <xmpDM:altTimecode
       xmpDM:timeValue="00:00:00:00"
       xmpDM:timeFormat="2997NonDropTimecode"/>
      <xmpDM:duration
       xmpDM:value="7800"
       xmpDM:scale="1/2997"/>
      <xmpDM:videoFrameSize
       stDim:w="720"
       stDim:h="480"
       stDim:unit="pixel"/>
      <xmpMM:History>
       <rdf:Seq>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:65294BA351EDE111BA25D0E0299F7394"
         stEvt:when="2012-08-23T20:40:31+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/metadata;/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:C2F75C0552EDE111BA25D0E0299F7394"
         stEvt:when="2012-08-23T20:40:41+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/metadata"/>
       </rdf:Seq>
      </xmpMM:History>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:C3355C2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">bg2.jpg</rdf:li>
       </rdf:Alt>
      </dc:title>
      <xmpMM:Ingredients>
       <rdf:Bag>
        <rdf:li
         stRef:instanceID="xmp.iid:5933D21264F0E11180B8D0D5AB388124"
         stRef:documentID="xmp.did:5833D21264F0E11180B8D0D5AB388124"
         stRef:fromPart="time:0d128000f25600"
         stRef:toPart="time:0d128000f25600"
         stRef:maskMarkers="All"/>
       </rdf:Bag>
      </xmpMM:Ingredients>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.comp"
       xmpMM:InstanceID="xmp.iid:C4355C2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">Comp 4</rdf:li>
       </rdf:Alt>
      </dc:title>
      <xmpMM:Ingredients>
       <rdf:Bag>
        <rdf:li
         stRef:instanceID="xmp.iid:519DBE7952F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d128000f25600"
         stRef:toPart="time:0d128000f25600"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:529DBE7952F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d128000f25600"
         stRef:toPart="time:0d128000f25600"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:539DBE7952F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d128000f25600"
         stRef:toPart="time:0d128000f25600"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:C0355C2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d128000f25600"
         stRef:toPart="time:0d128000f25600"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:C1355C2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d128000f25600"
         stRef:toPart="time:0d128000f25600"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:C2355C2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d128000f25600"
         stRef:toPart="time:0d128000f25600"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:C3355C2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d128000f25600"
         stRef:toPart="time:0d128000f25600"
         stRef:maskMarkers="None"/>
       </rdf:Bag>
      </xmpMM:Ingredients>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:C5355C2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">Comp 4</rdf:li>
       </rdf:Alt>
      </dc:title>
      <xmpMM:Ingredients>
       <rdf:Bag>
        <rdf:li
         stRef:instanceID="xmp.iid:C4355C2F53F4E11188E1CBA262DF7ED4"
         stRef:fromPart="time:0d128000f25600"
         stRef:toPart="time:0d128000f25600"
         stRef:maskMarkers="None"/>
       </rdf:Bag>
      </xmpMM:Ingredients>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:C6355C2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">Black Solid 3</rdf:li>
       </rdf:Alt>
      </dc:title>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:C7355C2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">Moja pesem</rdf:li>
       </rdf:Alt>
      </dc:title>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:C8355C2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">ALYA</rdf:li>
       </rdf:Alt>
      </dc:title>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:C9355C2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">White Solid 2</rdf:li>
       </rdf:Alt>
      </dc:title>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       dc:format="application/vnd.adobe.aftereffects.layer"
       xmpMM:InstanceID="xmp.iid:CA355C2F53F4E11188E1CBA262DF7ED4">
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">bg2.jpg</rdf:li>
       </rdf:Alt>
      </dc:title>
      <xmpMM:Ingredients>
       <rdf:Bag>
        <rdf:li
         stRef:instanceID="xmp.iid:5933D21264F0E11180B8D0D5AB388124"
         stRef:documentID="xmp.did:5833D21264F0E11180B8D0D5AB388124"
         stRef:fromPart="time:0d76800f25600"
         stRef:toPart="time:0d76800f25600"
         stRef:maskMarkers="All"/>
       </rdf:Bag>
      </xmpMM:Ingredients>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       xmpDM:trackNumber="1"
       xmpDM:logComment="www.dvdvideosoft.com"
       xmpDM:audioSampleRate="48000"
       xmpDM:audioSampleType="Compressed"
       xmpDM:audioChannelType="Stereo"
       xmpDM:partOfCompilation="false"
       xmp:CreateDate="2012"
       xmp:MetadataDate="2012-07-29T18:14:46+02:00"
       xmp:ModifyDate="2012-07-29T18:14:42+02:00"
       xmpMM:InstanceID="xmp.iid:EFCDFCEC97D9E111BF23A9FF07831E55"
       xmpMM:DocumentID="xmp.did:EDCDFCEC97D9E111BF23A9FF07831E55"
       xmpMM:OriginalDocumentID="xmp.did:EDCDFCEC97D9E111BF23A9FF07831E55">
      <xmpMM:History>
       <rdf:Seq>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:EECDFCEC97D9E111BF23A9FF07831E55"
         stEvt:when="2012-07-29T18:14:42+02:00"
         stEvt:softwareAgent="Adobe Premiere Pro 5.5"
         stEvt:changed="/metadata;/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:EFCDFCEC97D9E111BF23A9FF07831E55"
         stEvt:when="2012-07-29T18:14:46+02:00"
         stEvt:softwareAgent="Adobe Premiere Pro 5.5"
         stEvt:changed="/metadata"/>
       </rdf:Seq>
      </xmpMM:History>
      <dc:title>
       <rdf:Alt>
        <rdf:li xml:lang="x-default">Alya - Moja pesem</rdf:li>
       </rdf:Alt>
      </dc:title>
      </rdf:Description>
     </rdf:li>
     <rdf:li>
      <rdf:Description
       xmp:CreateDate="2012-08-24T18:38:48Z"
       xmp:ModifyDate="2012-08-24T18:39:55Z"
       xmp:CreatorTool="Adobe Premiere Pro 5.5"
       xmp:MetadataDate="2012-09-01T18:53:49+02:00"
       xmpDM:videoFrameRate="25.000000"
       xmpDM:videoFieldOrder="Progressive"
       xmpDM:videoPixelAspectRatio="17280/17280"
       xmpDM:audioSampleRate="48000"
       xmpDM:audioSampleType="Compressed"
       xmpDM:audioChannelType="Stereo"
       xmpDM:startTimeScale="25"
       xmpDM:startTimeSampleSize="1"
       xmpDM:videoAlphaMode="none"
       xmpMM:OriginalDocumentID="xmp.did:D1BEB3351AEEE1119BFCD6521D1187C9"
       xmpMM:InstanceID="xmp.iid:FE4E80EA19EEE11180F5C2DBA05DF1FB"
       xmpMM:DocumentID="xmp.did:D3BEB3351AEEE1119BFCD6521D1187C9"
       dc:format="H.264">
      <xmpDM:projectRef
       xmpDM:type="movie"/>
      <xmpDM:videoFrameSize
       stDim:w="1920"
       stDim:h="1080"
       stDim:unit="pixel"/>
      <xmpDM:startTimecode
       xmpDM:timeFormat="25Timecode"
       xmpDM:timeValue="00:00:00:00"/>
      <xmpDM:altTimecode
       xmpDM:timeValue="00:00:00:00"
       xmpDM:timeFormat="25Timecode"/>
      <xmpDM:duration
       xmpDM:value="3878400"
       xmpDM:scale="1/90000"/>
      <xmpDM:Tracks>
       <rdf:Bag>
        <rdf:li>
         <rdf:Description
          xmpDM:trackName="Markers"
          xmpDM:frameRate="f254016000000">
         <xmpDM:markers>
          <rdf:Seq>
           <rdf:li
            xmpDM:startTime="2997388800000"
            xmpDM:duration="10160640000"/>
           <rdf:li
            xmpDM:startTime="9185218560000"
            xmpDM:duration="10160640000"/>
           <rdf:li
            xmpDM:startTime="9327467520000"
            xmpDM:duration="10160640000"/>
           <rdf:li
            xmpDM:startTime="10231764480000"
            xmpDM:duration="10160640000"/>
          </rdf:Seq>
         </xmpDM:markers>
         </rdf:Description>
        </rdf:li>
       </rdf:Bag>
      </xmpDM:Tracks>
      <xmpMM:Ingredients>
       <rdf:Bag>
        <rdf:li
         stRef:instanceID="xmp.iid:1705054974ECE111B9A38987502929D1"
         stRef:documentID="xmp.did:5855F98372ECE111B9A38987502929D1"
         stRef:fromPart="time:9947266560000f254016000000d1961003520000f254016000000"
         stRef:toPart="time:8128512000000f254016000000d1961003520000f254016000000"
         stRef:filePath="MVI_9186.MOV"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:1705054974ECE111B9A38987502929D1"
         stRef:documentID="xmp.did:5855F98372ECE111B9A38987502929D1"
         stRef:fromPart="time:4064256000000f254016000000d2550320640000f254016000000"
         stRef:toPart="time:8382528000000f254016000000d2550320640000f254016000000"
         stRef:filePath="MVI_9186.MOV"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:1705054974ECE111B9A38987502929D1"
         stRef:documentID="xmp.did:5855F98372ECE111B9A38987502929D1"
         stRef:fromPart="time:6472327680000f254016000000d660441600000f254016000000"
         stRef:toPart="time:3922007040000f254016000000d660441600000f254016000000"
         stRef:filePath="MVI_9186.MOV"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:3C5F4D8EE8E9E111B624CD95E5AA602F"
         stRef:documentID="xmp.did:395F4D8EE8E9E111B624CD95E5AA602F"
         stRef:fromPart="time:12497587200000f254016000000d1473292800000f254016000000"
         stRef:toPart="time:5334336000000f254016000000d1473292800000f254016000000"
         stRef:filePath="MVI_9190.MOV"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:3C5F4D8EE8E9E111B624CD95E5AA602F"
         stRef:documentID="xmp.did:395F4D8EE8E9E111B624CD95E5AA602F"
         stRef:fromPart="time:15515297280000f254016000000d2499517440000f254016000000"
         stRef:toPart="time:8352046080000f254016000000d2499517440000f254016000000"
         stRef:filePath="MVI_9190.MOV"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:3C5F4D8EE8E9E111B624CD95E5AA602F"
         stRef:documentID="xmp.did:395F4D8EE8E9E111B624CD95E5AA602F"
         stRef:fromPart="time:14113128960000f254016000000d863654400000f254016000000"
         stRef:toPart="time:9368110080000f254016000000d863654400000f254016000000"
         stRef:filePath="MVI_9190.MOV"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:3D5F4D8EE8E9E111B624CD95E5AA602F"
         stRef:documentID="xmp.did:355F4D8EE8E9E111B624CD95E5AA602F"
         stRef:fromPart="time:25096780800000f254016000000d843333120000f254016000000"
         stRef:toPart="time:10089515520000f254016000000d843333120000f254016000000"
         stRef:filePath="MVI_9164.MOV"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:3D5F4D8EE8E9E111B624CD95E5AA602F"
         stRef:documentID="xmp.did:355F4D8EE8E9E111B624CD95E5AA602F"
         stRef:fromPart="time:17679513600000f254016000000d5049838080000f254016000000"
         stRef:toPart="time:5334336000000f254016000000d5049838080000f254016000000"
         stRef:filePath="MVI_9164.MOV"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:3D5F4D8EE8E9E111B624CD95E5AA602F"
         stRef:documentID="xmp.did:355F4D8EE8E9E111B624CD95E5AA602F"
         stRef:fromPart="time:17679513600000f254016000000d5517227520000f254016000000"
         stRef:toPart="time:5334336000000f254016000000d5517227520000f254016000000"
         stRef:filePath="MVI_9164.MOV"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:3D5F4D8EE8E9E111B624CD95E5AA602F"
         stRef:documentID="xmp.did:355F4D8EE8E9E111B624CD95E5AA602F"
         stRef:fromPart="time:12863370240000f254016000000d4064256000000f254016000000"
         stRef:toPart="time:518192640000f254016000000d4064256000000f254016000000"
         stRef:filePath="MVI_9164.MOV"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:3D5F4D8EE8E9E111B624CD95E5AA602F"
         stRef:documentID="xmp.did:355F4D8EE8E9E111B624CD95E5AA602F"
         stRef:fromPart="time:12863370240000f254016000000d4064256000000f254016000000"
         stRef:toPart="time:518192640000f254016000000d4064256000000f254016000000"
         stRef:filePath="MVI_9164.MOV"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:45F9717706EEE1119BFCD6521D1187C9"
         stRef:documentID="xmp.did:43F9717706EEE1119BFCD6521D1187C9"
         stRef:fromPart="time:11735539200000f254016000000d1341204480000f254016000000"
         stRef:toPart="time:6787307520000f254016000000d1341204480000f254016000000"
         stRef:filePath="MVI_9177.MOV"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:5255F98372ECE111B9A38987502929D1"
         stRef:documentID="xmp.did:5055F98372ECE111B9A38987502929D1"
         stRef:fromPart="time:14661803520000f254016000000d467389440000f254016000000"
         stRef:toPart="time:2997388800000f254016000000d467389440000f254016000000"
         stRef:filePath="MVI_9179.MOV"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:5555F98372ECE111B9A38987502929D1"
         stRef:documentID="xmp.did:5355F98372ECE111B9A38987502929D1"
         stRef:fromPart="time:18543168000000f254016000000d457228800000f254016000000"
         stRef:toPart="time:3464778240000f254016000000d457228800000f254016000000"
         stRef:filePath="MVI_9180.MOV"
         stRef:maskMarkers="None"/>
        <rdf:li
         stRef:instanceID="xmp.iid:EFCDFCEC97D9E111BF23A9FF07831E55"
         stRef:documentID="xmp.did:EDCDFCEC97D9E111BF23A9FF07831E55"
         stRef:fromPart="time:12589032960000f254016000000d10932848640000f254016000000"
         stRef:toPart="time:0d10932848640000f254016000000"
         stRef:filePath="Alya - Moja pesem.mp3"
         stRef:maskMarkers="None"/>
       </rdf:Bag>
      </xmpMM:Ingredients>
      <xmpMM:History>
       <rdf:Seq>
        <rdf:li
         stEvt:action="created"
         stEvt:instanceID="xmp.iid:D1BEB3351AEEE1119BFCD6521D1187C9"
         stEvt:when="2012-08-24T20:35:07+02:00"
         stEvt:softwareAgent="Adobe Premiere Pro 5.5"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:D2BEB3351AEEE1119BFCD6521D1187C9"
         stEvt:when="2012-08-24T20:38:50+02:00"
         stEvt:softwareAgent="Adobe Premiere Pro 5.5"
         stEvt:changed="/metadata"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:D3BEB3351AEEE1119BFCD6521D1187C9"
         stEvt:when="2012-08-24T20:38:50+02:00"
         stEvt:softwareAgent="Adobe Premiere Pro 5.5"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:FD4E80EA19EEE11180F5C2DBA05DF1FB"
         stEvt:when="2012-08-24T20:39:55+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/metadata;/content"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:FE4E80EA19EEE11180F5C2DBA05DF1FB"
         stEvt:when="2012-08-24T20:40:05+02:00"
         stEvt:softwareAgent="Adobe After Effects CS5.5 10.5 Windows 64"
         stEvt:changed="/metadata"/>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="xmp.iid:320FA19553F4E1118E5EF593D63F81DA"
         stEvt:when="2012-09-01T18:53:49+02:00"
         stEvt:softwareAgent="Adobe Premiere Pro 5.5"
         stEvt:changed="/metadata"/>
       </rdf:Seq>
      </xmpMM:History>
      <creatorAtom:windowsAtom
       creatorAtom:extension=".prproj"
       creatorAtom:invocationFlags="/L"
       creatorAtom:uncProjectPath="\\?\E:\alyaspot\Posnetki\Nova mapa\alya_trailer.prproj"/>
      <creatorAtom:macAtom
       creatorAtom:applicationCode="1347449455"
       creatorAtom:invocationAppleEvent="1129468018"/>
      </rdf:Description>
     </rdf:li>
    </rdf:Bag>
   </xmpMM:Pantry>
  </rdf:Description>
 </rdf:RDF>
</x:xmpmeta>





















<?xpacket end="w"?>`

const data_5kb = `<?xpacket begin="" id="W5M0MpCehiHzreSzNTczkc9d"?>
<x:xmpmeta xmlns:x="adobe:ns:meta/" x:xmptk="Adobe XMP Core 5.5-c021 79.155241, 2013/11/25-21:10:40        ">
 <rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#">
  <rdf:Description rdf:about=""
    xmlns:xmp="http://ns.adobe.com/xap/1.0/"
    xmlns:xmpDM="http://ns.adobe.com/xmp/1.0/DynamicMedia/"
    xmlns:stDim="http://ns.adobe.com/xap/1.0/sType/Dimensions#"
    xmlns:xmpMM="http://ns.adobe.com/xap/1.0/mm/"
    xmlns:stEvt="http://ns.adobe.com/xap/1.0/sType/ResourceEvent#"
    xmlns:stRef="http://ns.adobe.com/xap/1.0/sType/ResourceRef#"
    xmlns:dc="http://purl.org/dc/elements/1.1/"
    xmlns:creatorAtom="http://ns.adobe.com/creatorAtom/1.0/"
   xmp:CreateDate="2015-03-23T22:18:20-04:00"
   xmp:ModifyDate="2015-03-23T22:18:29-04:00"
   xmp:MetadataDate="2015-03-23T22:18:29-04:00"
   xmp:CreatorTool="Adobe Premiere Pro CC (Macintosh)"
   xmpDM:startTimeScale="24"
   xmpDM:startTimeSampleSize="1"
   xmpDM:videoFrameRate="24.000000"
   xmpDM:videoFieldOrder="Progressive"
   xmpDM:videoPixelAspectRatio="1/1"
   xmpDM:audioSampleRate="48000"
   xmpDM:audioSampleType="16Int"
   xmpDM:audioChannelType="Stereo"
   xmpMM:InstanceID="xmp.iid:5d95d24a-457a-44bd-a92c-0786456548fd"
   xmpMM:DocumentID="009fd353-c6ed-1e3f-5ae1-c92200000053"
   xmpMM:OriginalDocumentID="xmp.did:d1ae4cc5-a458-4eb7-9394-30ec546e2528"
   dc:format="QuickTime">
   <xmpDM:duration
    xmpDM:value="203"
    xmpDM:scale="1/24"/>
   <xmpDM:altTimecode
    xmpDM:timeValue="01:02:03:04"
    xmpDM:timeFormat="24Timecode"/>
   <xmpDM:videoFrameSize
    stDim:w="1280"
    stDim:h="720"
    stDim:unit="pixel"/>
   <xmpDM:startTimecode
    xmpDM:timeFormat="24Timecode"
    xmpDM:timeValue="00:00:00:00"/>
   <xmpDM:projectRef
    xmpDM:type="movie"/>
   <xmpMM:History>
    <rdf:Seq>
     <rdf:li
      stEvt:action="saved"
      stEvt:instanceID="a985eefc-d371-6b9c-d554-0f0200000080"
      stEvt:when="2015-03-23T22:18:29-04:00"
      stEvt:softwareAgent="Adobe Premiere Pro CC (Macintosh)"
      stEvt:changed="/"/>
     <rdf:li
      stEvt:action="created"
      stEvt:instanceID="xmp.iid:8c22b83c-cfd3-43fe-b402-2d96d4a3d4d4"
      stEvt:when="2015-03-23T22:18:20-04:00"
      stEvt:softwareAgent="Adobe Premiere Pro CC (Macintosh)"/>
     <rdf:li
      stEvt:action="saved"
      stEvt:instanceID="xmp.iid:3f3bd4e9-9370-4b25-bb36-ac7d7184bbd5"
      stEvt:when="2015-03-23T22:18:29-04:00"
      stEvt:softwareAgent="Adobe Premiere Pro CC (Macintosh)"
      stEvt:changed="/"/>
     <rdf:li
      stEvt:action="saved"
      stEvt:instanceID="xmp.iid:5d95d24a-457a-44bd-a92c-0786456548fd"
      stEvt:when="2015-03-23T22:18:29-04:00"
      stEvt:softwareAgent="Adobe Premiere Pro CC (Macintosh)"
      stEvt:changed="/metadata"/>
    </rdf:Seq>
   </xmpMM:History>
   <xmpMM:Ingredients>
    <rdf:Bag>
     <rdf:li
      stRef:instanceID="47a9a75b-3a6b-0990-173c-d5a900000074"
      stRef:documentID="aeffd160-efa3-71bf-610a-09bd00000047"
      stRef:fromPart="time:0d2154055680000f254016000000"
      stRef:toPart="time:0d2154055680000f254016000000"
      stRef:filePath="CornerKick.mov"
      stRef:maskMarkers="None"/>
     <rdf:li
      stRef:instanceID="47a9a75b-3a6b-0990-173c-d5a900000074"
      stRef:documentID="aeffd160-efa3-71bf-610a-09bd00000047"
      stRef:fromPart="time:0d2154055680000f254016000000"
      stRef:toPart="time:0d2154055680000f254016000000"
      stRef:filePath="CornerKick.mov"
      stRef:maskMarkers="None"/>
    </rdf:Bag>
   </xmpMM:Ingredients>
   <xmpMM:Pantry>
    <rdf:Bag>
     <rdf:li>
      <rdf:Description
       xmp:CreateDate="2014-03-20T21:34:41Z"
       xmp:ModifyDate="2015-03-23T22:10:43-04:00"
       xmp:MetadataDate="2015-03-23T22:10:43-04:00"
       xmpDM:startTimeScale="5000"
       xmpDM:startTimeSampleSize="200"
       xmpMM:InstanceID="47a9a75b-3a6b-0990-173c-d5a900000074"
       xmpMM:DocumentID="aeffd160-efa3-71bf-610a-09bd00000047"
       xmpMM:OriginalDocumentID="xmp.did:cd1f24ab-d821-4bad-aa79-bdc9baaabe14">
      <xmpDM:duration
       xmpDM:value="21200"
       xmpDM:scale="1/2500"/>
      <xmpDM:altTimecode
       xmpDM:timeValue="00:00:00:00"
       xmpDM:timeFormat="25Timecode"/>
      <xmpMM:History>
       <rdf:Seq>
        <rdf:li
         stEvt:action="saved"
         stEvt:instanceID="47a9a75b-3a6b-0990-173c-d5a900000074"
         stEvt:when="2015-03-23T22:10:43-04:00"
         stEvt:softwareAgent="Adobe Premiere Pro CC (Macintosh)"
         stEvt:changed="/"/>
       </rdf:Seq>
      </xmpMM:History>
      </rdf:Description>
     </rdf:li>
    </rdf:Bag>
   </xmpMM:Pantry>
   <xmpMM:DerivedFrom
    stRef:instanceID="xmp.iid:8c22b83c-cfd3-43fe-b402-2d96d4a3d4d4"
    stRef:documentID="xmp.did:8c22b83c-cfd3-43fe-b402-2d96d4a3d4d4"
    stRef:originalDocumentID="xmp.did:8c22b83c-cfd3-43fe-b402-2d96d4a3d4d4"/>
   <creatorAtom:windowsAtom
    creatorAtom:extension=".prproj"
    creatorAtom:invocationFlags="/L"/>
   <creatorAtom:macAtom
    creatorAtom:applicationCode="1347449455"
    creatorAtom:invocationAppleEvent="1129468018"
    creatorAtom:posixProjectPath="/Users/alandabul/Documents/Adobe/Premiere Pro/7.0/Untitled.prproj"/>
  </rdf:Description>
 </rdf:RDF>
</x:xmpmeta>





















<?xpacket end="w"?>`
