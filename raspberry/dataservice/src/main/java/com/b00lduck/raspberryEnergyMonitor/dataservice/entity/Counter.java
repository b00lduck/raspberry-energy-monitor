package com.b00lduck.raspberryEnergyMonitor.dataservice.entity;

import javax.persistence.*;
import java.util.List;

/**
 * @author Daniel Zerlett (daniel@zerlett.eu)
 * @created 24.07.2015
 */
@Entity
public class Counter {

	@Id
	@GeneratedValue(strategy = GenerationType.TABLE)
	private Long id;

	private String name;

	private String unit;

	@OneToMany(mappedBy = "counter")
	private List<CounterEvent> counterEvents;

    public Long getId() {
		return id;
	}

	public String getName() {
		return name;
	}

	public void setName(final String name) {
		this.name = name;
	}

	public String getUnit() {
		return unit;
	}

	public void setUnit(final String unit) {
		this.unit = unit;
	}

}
