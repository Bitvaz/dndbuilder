<script>
  document.getElementById('class-select').addEventListener('change', function() {
    const selectedClass = this.value;

    const stats = {
      fighter: {
        strength: 15,
        vigor: 15,
        agility: 15,
        dexterity: 15,
        will: 15,
        knowledge: 15,
        resourcefulness: 15
      },
      mage: {
        strength: 10,
        vigor: 12,
        agility: 13,
        dexterity: 14,
        will: 18,
        knowledge: 20,
        resourcefulness: 16
      },
      rogue: {
        strength: 12,
        vigor: 14,
        agility: 18,
        dexterity: 17,
        will: 13,
        knowledge: 12,
        resourcefulness: 16
      }
    };

    if (stats[selectedClass]) {
      document.getElementById('strength').value = stats[selectedClass].strength;
      document.getElementById('vigor').value = stats[selectedClass].vigor;
      document.getElementById('agility').value = stats[selectedClass].agility;
      document.getElementById('dexterity').value = stats[selectedClass].dexterity;
      document.getElementById('will').value = stats[selectedClass].will;
      document.getElementById('knowledge').value = stats[selectedClass].knowledge;
      document.getElementById('resourcefulness').value = stats[selectedClass].resourcefulness;
    }
  });
</script>