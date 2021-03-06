/*
Copyright (C) 2017 Verizon. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ast

type FieldDescription struct {
	FieldName string
	FieldType string
}

type TableDescription struct {
	Fields     []*FieldDescription
	PrimaryKey *PrimaryKey
}

type PrimaryKey struct {
	Partitioning []string
	Clustering   []string
}

type Order struct {
	FieldName string
	Ascending bool
}

type Property interface {
	PropertyMarker()
}

type ClusteringOrder struct {
	Order []*Order
}

func (co *ClusteringOrder) PropertyMarker() {}

type CompactStorage struct{}

func (cs *CompactStorage) PropertyMarker() {}
