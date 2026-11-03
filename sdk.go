package sdk

import "time"

type NavItem struct {
	ID, Label, Href string
	Icon            string // svg path markup
	Status          string // "ok" | "warn" | ""
}

type Row struct {
	Name, Value string
	Percent     int
	Warn        bool
}

type Widget struct {
	Title, Badge string
	Warn         bool
	Rows         []Row
}

type Event struct {
	ID, Title, Tag, Source string
	Start, End             time.Time
}

type Alert struct {
	ID, Module, Title, Body string
	Severity                string // "info" | "warn" | "critical"
	DedupeKey               string
	Link                    string
}